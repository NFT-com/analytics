import * as pulumi from '@pulumi/pulumi';
import * as aws from '@pulumi/aws';
import { getResourceName, getTags } from '../helper'
import { SharedInfraOutput } from '../defs'

const event_db = `host=${process.env.DB_EVENT_HOST} port=${process.env.DB_PORT} user=${process.env.DB_USER} password=${process.env.DB_PASSWORD} dbname=${process.env.DB_NAME}`
const graph_db = `host=${process.env.DB_GRAPH_HOST} port=${process.env.DB_PORT} user=${process.env.DB_USER} password=${process.env.DB_PASSWORD} dbname=${process.env.DB_NAME}`
const execRole = 'arn:aws:iam::016437323894:role/ecsTaskExecutionRole'
const taskRole = 'arn:aws:iam::016437323894:role/ECSServiceTask'

const tags = {
  service: 'analytics',
}

/*

const attachLBListeners = (
  lb: aws.lb.LoadBalancer,
  tg: aws.lb.TargetGroup,
): void => {
  new aws.lb.Listener('listener_http_dev_gql_ecs', {
    defaultActions: [
      {
        order: 1,
        redirect: {
          port: '443',
          protocol: 'HTTPS',
          statusCode: 'HTTP_301',
        },
        type: 'redirect',
      },
    ],
    loadBalancerArn: lb.arn,
    port: 80,
    protocol: 'HTTP',
    tags: getTags(tags),
  })

  new aws.lb.Listener('listener_https_dev_gql_ecs', {
    certificateArn:
      'arn:aws:acm:us-east-1:016437323894:certificate/0c01a3a8-59c4-463a-87ec-5c487695f09e',
    defaultActions: [
      {
        targetGroupArn: tg.arn,
        type: 'forward',
      },
    ],
    loadBalancerArn: lb.arn,
    port: 443,
    protocol: 'HTTPS',
    sslPolicy: 'ELBSecurityPolicy-2016-08',
    tags: getTags(tags),
  })
}

const createEcsTargetGroup = (
  infraOutput: SharedInfraOutput,
): aws.lb.TargetGroup => {
  return new aws.lb.TargetGroup('tg_gql_ecs', {
    healthCheck: {
      interval: 15,
      matcher: '200-399',
      path: '/.well-known/apollo/server-health',
      timeout: 5,
      unhealthyThreshold: 5,
    },
    name: getResourceName('gql-ecs'),
    port: 8080,
    protocol: 'HTTP',
    protocolVersion: 'HTTP1',
    stickiness: {
      enabled: false,
      type: 'lb_cookie',
    },
    targetType: 'ip',
    vpcId: infraOutput.vpcId,
    tags: getTags(tags),
  })
}

const createEcsLoadBalancer = (
  infraOutput: SharedInfraOutput,
): aws.lb.LoadBalancer => {
  return new aws.lb.LoadBalancer('lb_gql_ecs', {
    ipAddressType: 'ipv4',
    name: getResourceName('gql-ecs'),
    securityGroups: [infraOutput.webSGId],
    subnets: infraOutput.publicSubnets,
    tags: getTags(tags),
  })
}
*/

export const createGraphTaskDefinition = (
    infraOutput: SharedInfraOutput,
): aws.ecs.TaskDefinition => {
    const resourceName = getResourceName('analytics-td-graph-api')
    const ecrImage = `${process.env.ECR_REGISTRY}/${infraOutput.analyticECRRepo}:graph-api`
    
    return new aws.ecs.TaskDefinition(resourceName, 
    {
        containerDefinitions: JSON.stringify([
            {
                command: ['--database',graph_db,'--enable-playground','-db-connection-limit','70','-l',process.env.LOG_LEVEL,'--enable-query-logging','--search-limit','20'],
                cpu: 0,
                entryPoint: ['/api'],
                essential: true,
                image: ecrImage,
                links: [],
                memoryReservation: 512,
                mountPoints: [],
                name: resourceName,
                portMappings: [
                    { 
                        containerPort: 8083,
                        hostPort: 8080,
                        protocol: 'tcp'
                    }
                ],                environment: [],
                volumesFrom: []
        }]),
        executionRoleArn: execRole,
        family: resourceName,
        cpu: '256',
        memory: '512',
        requiresCompatibilities: ['EC2'],
        taskRoleArn: taskRole,
    })
}

export const createAggregationTaskDefinition = (
    infraOutput: SharedInfraOutput,
): aws.ecs.TaskDefinition => {
    const resourceName = getResourceName('analytics-td-aggregation-api')
    const ecrImage = `${process.env.ECR_REGISTRY}/${infraOutput.analyticECRRepo}:aggregation-api`

    return new aws.ecs.TaskDefinition(resourceName, 
    {
        containerDefinitions: JSON.stringify([
            {
                command: ['--events-api','events-api:8080','-l',process.env.LOG_LEVEL],
                cpu: 0,
                entryPoint: ['/api'],
                essential: true,
                image: ecrImage,
                links: [],
                memoryReservation: 512,
                mountPoints: [],
                name: resourceName,
                portMappings: [
                    { 
                        containerPort: 8084,
                        hostPort: 8080,
                        protocol: 'tcp'
                    }
                ],
                environment: [],
                volumesFrom:[],
        }]),
        executionRoleArn: execRole,
        family: resourceName,
        cpu: '256',
        memory: '512',
        requiresCompatibilities: ['EC2'],
        taskRoleArn: taskRole,
    })
}

export const createEventsTaskDefinition = (
    infraOutput: SharedInfraOutput,
): aws.ecs.TaskDefinition => {
    const resourceName = getResourceName('analytics-td-events-api')
    const ecrImage = `${process.env.ECR_REGISTRY}/${infraOutput.analyticECRRepo}:events-api`

    return new aws.ecs.TaskDefinition(resourceName, 
    {
        containerDefinitions: JSON.stringify([
            {
                command: ['--database',event_db,'-l',process.env.LOG_LEVEL,'--db-connection-limit','70','--batch-size','100','--enable-query-logging'],
                cpu: 0,
                entryPoint: ['/api'],
                environment: [],
                essential: true,
                image: ecrImage,
                links: [],
                memoryReservation: 512,
                mountPoints: [],
                name: resourceName,
                portMappings: [
                    { 
                        containerPort: 8085,
                        hostPort: 8080,
                        protocol: 'tcp'
                    }
                ],
                volumesFrom: []
        }]),
        executionRoleArn: execRole,
        family: resourceName,
        cpu: '256',
        memory: '512',
        requiresCompatibilities: ['EC2'],
        taskRoleArn: taskRole,
    })
}

const createEcsAsgLaunchConfig = (
    infraOutput: SharedInfraOutput,
): aws.ec2.LaunchConfiguration => {
    const launchConfigSG = infraOutput.webSGId
    const clusterName = getResourceName('analytics')
    const resourceName = getResourceName('analytics-asg-launchconfig')
    const ec2UserData =
    `#!/bin/bash
    echo ECS_CLUSTER=${clusterName} >> /etc/ecs/ecs.config
    echo ECS_BACKEND_HOST= >> /etc/ecs/ecs.config`

    return new aws.ec2.LaunchConfiguration(resourceName, {
        associatePublicIpAddress: true,
        iamInstanceProfile: 'arn:aws:iam::016437323894:instance-profile/ecsInstanceRole',
        imageId: 'ami-0f863d7367abe5d6f',  //latest amzn linux 2 ecs-optimized ami in us-east-1
        instanceType: 't3.medium', // upgrade when ready 
        keyName: 'indexer_dev_key', // same key for both indexer/analytics instances
        name: resourceName,
        rootBlockDevice: {
            deleteOnTermination: false,
            volumeSize: 30,
            volumeType: 'gp2',
        },
        securityGroups: [launchConfigSG],
        userData: ec2UserData,
    })
}

const createEcsASG = (
    config: pulumi.Config,
    infraOutput: SharedInfraOutput,
): aws.autoscaling.Group => {
    const resourceName = getResourceName('analytics-asg')
    return new aws.autoscaling.Group(resourceName, {
        defaultCooldown: 300,
        desiredCapacity: 1,
        healthCheckGracePeriod: 0,
        healthCheckType: 'EC2',
        launchConfiguration: createEcsAsgLaunchConfig(infraOutput),
        maxSize: 1,
        minSize: 1,
        name: resourceName,
        serviceLinkedRoleArn: 'arn:aws:iam::016437323894:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling',
        tags: [
            {
                key: 'Description',
                propagateAtLaunch: true,
                value: 'This instance is the part of the Auto Scaling group which was created through ECS Console',
            },
            {
                key: 'AmazonECSManaged',
                propagateAtLaunch: true,
                value: '',
            },
            {
                key: 'Name',
                propagateAtLaunch: true,
                value: resourceName,
            },
        ],
        vpcZoneIdentifiers: infraOutput.publicSubnets,
    })
}

const createEcsCapacityProvider = (
    config: pulumi.Config,
    infraOutput: SharedInfraOutput,
): aws.ecs.CapacityProvider => {
    const resourceName = getResourceName('analytics-cp')
    const { arn: arn_asg } = createEcsASG(config, infraOutput)
    return new aws.ecs.CapacityProvider(resourceName, {
        autoScalingGroupProvider: {
            autoScalingGroupArn: arn_asg,
            managedScaling: {
                instanceWarmupPeriod: 300,
                maximumScalingStepSize: 1,
                minimumScalingStepSize: 1,
                status: 'DISABLED',
                targetCapacity: 100,
            },
            managedTerminationProtection: 'DISABLED',
        },
        name: resourceName,
    })
}

export const createEcsCluster = (
    config: pulumi.Config,
    infraOutput: SharedInfraOutput,
): aws.ecs.Cluster => {
    const resourceName = getResourceName('analytics')
    const { name: capacityProvider } = createEcsCapacityProvider(config, infraOutput)
    const cluster = new aws.ecs.Cluster(resourceName, 
    {
        name: resourceName,
        settings: [
            {
            name: 'containerInsights',
            value: 'enabled',
        }],
        capacityProviders: [capacityProvider]
    })

    return cluster 
}