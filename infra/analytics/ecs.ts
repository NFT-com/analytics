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

const applyEventServiceAutoscaling = (
  config: pulumi.Config,
  service: aws.ecs.Service,
): void => {
  const target = new aws.appautoscaling.Target('analytics-svcAsg-target', {
    maxCapacity: 1,
    minCapacity: 1,
    resourceId: service.id.apply((id) => id.split(':').pop() || ''),
    scalableDimension: 'ecs:service:DesiredCount',
    serviceNamespace: 'ecs',
  })

  new aws.appautoscaling.Policy('analytics-svcAsg-policy', {
    policyType: 'TargetTrackingScaling',
    resourceId: target.resourceId,
    scalableDimension: target.scalableDimension,
    serviceNamespace: target.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
      targetValue: 60,
      predefinedMetricSpecification: {
        predefinedMetricType: 'ECSServiceAverageCPUUtilization',
      },
      scaleInCooldown: 360,
    },
  })
}

const applyGraphServiceAutoscaling = (
  config: pulumi.Config,
  service: aws.ecs.Service,
): void => {
  const target = new aws.appautoscaling.Target('analytics-svcAsg-graph-target', {
    maxCapacity: 1,
    minCapacity: 1,
    resourceId: service.id.apply((id) => id.split(':').pop() || ''),
    scalableDimension: 'ecs:service:DesiredCount',
    serviceNamespace: 'ecs',
  })

  new aws.appautoscaling.Policy('analytics-svcAsg-graph-policy', {
    policyType: 'TargetTrackingScaling',
    resourceId: target.resourceId,
    scalableDimension: target.scalableDimension,
    serviceNamespace: target.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
      targetValue: 60,
      predefinedMetricSpecification: {
        predefinedMetricType: 'ECSServiceAverageCPUUtilization',
      },
      scaleInCooldown: 360,
    },
  })
}

const applyAggregationServiceAutoscaling = (
  config: pulumi.Config,
  service: aws.ecs.Service,
): void => {
  const target = new aws.appautoscaling.Target('analytics-svcAsg-aggregation-target', {
    maxCapacity: 1,
    minCapacity: 1,
    resourceId: service.id.apply((id) => id.split(':').pop() || ''),
    scalableDimension: 'ecs:service:DesiredCount',
    serviceNamespace: 'ecs',
  })

  new aws.appautoscaling.Policy('analytics-svcAsg-aggregation-policy', {
    policyType: 'TargetTrackingScaling',
    resourceId: target.resourceId,
    scalableDimension: target.scalableDimension,
    serviceNamespace: target.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
      targetValue: 60,
      predefinedMetricSpecification: {
        predefinedMetricType: 'ECSServiceAverageCPUUtilization',
      },
      scaleInCooldown: 360,
    },
  })
}

const attachEventLBListeners = (
  lb: aws.lb.LoadBalancer,
  tg: aws.lb.TargetGroup,
): void => {
  new aws.lb.Listener('analytics-listener-http', {
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

  new aws.lb.Listener('analytics-listener-https', {
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

const attachGraphLBListeners = (
  lb: aws.lb.LoadBalancer,
  tg: aws.lb.TargetGroup,
): void => {
  new aws.lb.Listener('analytics-graph-listener-http', {
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

  new aws.lb.Listener('analytics-graph-listener-https', {
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

const attachAggregationLBListeners = (
  lb: aws.lb.LoadBalancer,
  tg: aws.lb.TargetGroup,
): void => {
  new aws.lb.Listener('analytics-aggregation-listener-http', {
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

  new aws.lb.Listener('analytics-aggregation-listener-https', {
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

const createEventTargetGroup = (
  infraOutput: SharedInfraOutput,
): aws.lb.TargetGroup => {
  const resourceName = getResourceName('analytics-lb-tg')
  return new aws.lb.TargetGroup(resourceName, {
    healthCheck: {
      interval: 15,
      port: '8085',
      matcher: '200-399',
      path: '/health',
      timeout: 5,
      unhealthyThreshold: 5,
    },
    name: resourceName,
    port: 8085,
    protocol: 'HTTP',
    protocolVersion: 'HTTP1',
    stickiness: {
      enabled: false,
      type: 'lb_cookie',
    },
    targetType: 'instance',
    vpcId: infraOutput.vpcId,
    tags: getTags(tags),
  })
}

const createGraphTargetGroup = (
  infraOutput: SharedInfraOutput,
): aws.lb.TargetGroup => {
  const resourceName = getResourceName('analytics-graph-lb-tg')
  return new aws.lb.TargetGroup(resourceName, {
    healthCheck: {
      interval: 15,
      port: '8083',
      matcher: '200-399',
      path: '/health', 
      timeout: 5,
      unhealthyThreshold: 5,
    },
    name: resourceName,
    port: 8083,
    protocol: 'HTTP',
    protocolVersion: 'HTTP1',
    stickiness: {
      enabled: false,
      type: 'lb_cookie',
    },
    targetType: 'instance',
    vpcId: infraOutput.vpcId,
    tags: getTags(tags),
  })
}

const createAggregationTargetGroup = (
  infraOutput: SharedInfraOutput,
): aws.lb.TargetGroup => {
  const resourceName = getResourceName('analytics-aggregation-lb-tg')
  return new aws.lb.TargetGroup(resourceName, {
    healthCheck: {
      interval: 15,
      port: '8084',
      matcher: '200-399',
      path: '/health', 
      timeout: 5,
      unhealthyThreshold: 5,
    },
    name: resourceName,
    port: 8084,
    protocol: 'HTTP',
    protocolVersion: 'HTTP1',
    stickiness: {
      enabled: false,
      type: 'lb_cookie',
    },
    targetType: 'instance',
    vpcId: infraOutput.vpcId,
    tags: getTags(tags),
  })
}

const createEventLoadBalancer = (
  infraOutput: SharedInfraOutput,
): aws.lb.LoadBalancer => {
  const resourceName = getResourceName('analytics-lb')
  return new aws.lb.LoadBalancer(resourceName, {
    ipAddressType: 'ipv4',
    name: resourceName,
    securityGroups: [infraOutput.webSGId],
    subnets: infraOutput.publicSubnets,
    tags: getTags(tags),
  })
}

const createGraphLoadBalancer = (
  infraOutput: SharedInfraOutput,
): aws.lb.LoadBalancer => {
  const resourceName = getResourceName('analytics-graph-lb')
  return new aws.lb.LoadBalancer(resourceName, {
    ipAddressType: 'ipv4',
    name: resourceName,
    securityGroups: [infraOutput.webSGId],
    subnets: infraOutput.publicSubnets,
    tags: getTags(tags),
  })
}

const createAggregationLoadBalancer = (
  infraOutput: SharedInfraOutput,
): aws.lb.LoadBalancer => {
  const resourceName = getResourceName('analytics-aggregation-lb')
  return new aws.lb.LoadBalancer(resourceName, {
    ipAddressType: 'ipv4',
    name: resourceName,
    securityGroups: [infraOutput.webSGId],
    subnets: infraOutput.publicSubnets,
    tags: getTags(tags),
  })
}

export const createGraphTaskDefinition = (
    infraOutput: SharedInfraOutput,
): aws.ecs.TaskDefinition => {
    const resourceName = getResourceName('analytics-td-graph-api')
    const ecrImage = `${process.env.ECR_REGISTRY}/${infraOutput.analyticECRRepo}:graph-api-${process.env.GIT_SHA || 'latest'}`
    
    return new aws.ecs.TaskDefinition(resourceName, 
    {
        containerDefinitions: JSON.stringify([
            {
                command: ['--database',graph_db,'--enable-playground','--db-connection-limit','70','--log-level',process.env.ANALYTICS_LOG_LEVEL,'--enable-query-logging','--search-limit','20'],
                cpu: 0,
                logConfiguration: {
                  logDriver: 'awslogs',
                  options: {
                    'awslogs-create-group': 'True',
                    'awslogs-group': `/ecs/${process.env.STAGE}-analytics-graph-api`,
                    'awslogs-region': 'us-east-1',
                    'awslogs-stream-prefix': 'analytics',
                  },
                },
                entryPoint: ['/api'],
                essential: true,
                image: ecrImage,
                links: [],
                memoryReservation: 512,
                mountPoints: [],
                name: resourceName,
                portMappings: [
                    { 
                        containerPort: 8080,
                        hostPort: 8083,
                        protocol: 'tcp'
                    }
                ],
                environment: [],
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
  const ecrImage = `${process.env.ECR_REGISTRY}/${infraOutput.analyticECRRepo}:aggregation-api-${process.env.GIT_SHA || 'latest'}`

  return new aws.ecs.TaskDefinition(resourceName, 
  {
      containerDefinitions: JSON.stringify([
          {
              command: ['--events-database',event_db,'--graph-database',graph_db,'--log-level',process.env.ANALYTICS_LOG_LEVEL,'--events-db-connection-limit','70','--enable-query-logging'],
              cpu: 0,
              logConfiguration: {
                logDriver: 'awslogs',
                options: {
                  'awslogs-create-group': 'True',
                  'awslogs-group': `/ecs/${process.env.STAGE}-analytics-aggregation-api`,
                  'awslogs-region': 'us-east-1',
                  'awslogs-stream-prefix': 'analytics',
                },
              },
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
                      containerPort: 8080,
                      hostPort: 8084,
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

export const createEventsTaskDefinition = (
    infraOutput: SharedInfraOutput,
): aws.ecs.TaskDefinition => {
    const resourceName = getResourceName('analytics-td-events-api')
    const ecrImage = `${process.env.ECR_REGISTRY}/${infraOutput.analyticECRRepo}:events-api-${process.env.GIT_SHA || 'latest'}`

    return new aws.ecs.TaskDefinition(resourceName, 
    {
        containerDefinitions: JSON.stringify([
            {
                command: ['--database',event_db,'--log-level',process.env.ANALYTICS_LOG_LEVEL,'--db-connection-limit','70','--batch-size','100','--enable-query-logging'],
                cpu: 0,
                entryPoint: ['/api'],
                logConfiguration: {
                  logDriver: 'awslogs',
                  options: {
                    'awslogs-create-group': 'True',
                    'awslogs-group': `/ecs/${process.env.STAGE}-analytics-events-api`,
                    'awslogs-region': 'us-east-1',
                    'awslogs-stream-prefix': 'analytics',
                  },
                },
                environment: [],
                essential: true,
                image: ecrImage,
                links: [],
                memoryReservation: 512,
                mountPoints: [],
                name: resourceName,
                portMappings: [
                    { 
                        containerPort: 8080,
                        hostPort: 8085,
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
    const resourceName = getResourceName('analytics-ec2')
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
    eventTaskDefinition: aws.ecs.TaskDefinition,
    aggregationTaskDefinition: aws.ecs.TaskDefinition,
    graphTaskDefinition: aws.ecs.TaskDefinition,
): aws.ecs.Cluster => {
    const resourceName = getResourceName('analytics')
    const { name: capacityProvider } = createEcsCapacityProvider(config, infraOutput)

    // create ecs cluster
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

    // create ecs event service (lb, tg, listeners, svc)
    const eventTargetGroup = createEventTargetGroup(infraOutput)
    const eventloadBalancer = createEventLoadBalancer(infraOutput)
    attachEventLBListeners(eventloadBalancer, eventTargetGroup)

    const eventServiceResourceName = getResourceName('analytics-event-svc')
    const eventService = new aws.ecs.Service(eventServiceResourceName, {
      cluster: cluster.arn,
      deploymentCircuitBreaker: {
        enable: true,
        rollback: true,
      },
      desiredCount: 1,
      deploymentMaximumPercent: 100,
      deploymentMinimumHealthyPercent: 0,
      enableEcsManagedTags: true,
      forceNewDeployment: true,
      healthCheckGracePeriodSeconds: 20,
      launchType: 'EC2',
      loadBalancers: [
        {
          containerName: eventTaskDefinition.family,
          containerPort: 8080,
          targetGroupArn: eventTargetGroup.arn,
        },
      ],
      name: eventServiceResourceName,
      taskDefinition: eventTaskDefinition.arn,
      tags: getTags(tags),
    })
  
    applyEventServiceAutoscaling(config, eventService)

    // create ecs aggregation service (lb, tg, listeners, svc)
    const aggregationTargetGroup = createAggregationTargetGroup(infraOutput)
    const aggregationloadBalancer = createAggregationLoadBalancer(infraOutput)
    attachAggregationLBListeners(aggregationloadBalancer, aggregationTargetGroup)

    const aggregationServiceResourceName = getResourceName('analytics-aggregation-svc')
    const aggregationService = new aws.ecs.Service(aggregationServiceResourceName, {
      cluster: cluster.arn,
      deploymentCircuitBreaker: {
        enable: true,
        rollback: true,
      },
      desiredCount: 1,
      deploymentMaximumPercent: 100,
      deploymentMinimumHealthyPercent: 0,
      enableEcsManagedTags: true,
      forceNewDeployment: true,
      healthCheckGracePeriodSeconds: 20,
      launchType: 'EC2',
      loadBalancers: [
        {
          containerName: aggregationTaskDefinition.family,
          containerPort: 8080,
          targetGroupArn: aggregationTargetGroup.arn,
        },
      ],
      name: aggregationServiceResourceName,
      taskDefinition: aggregationTaskDefinition.arn,
      tags: getTags(tags),
    })
    applyAggregationServiceAutoscaling(config, aggregationService)

    // create ecs graph service (lb, tg, listeners, svc)
    const graphTargetGroup = createGraphTargetGroup(infraOutput)
    const graphloadBalancer = createGraphLoadBalancer(infraOutput)
    attachGraphLBListeners(graphloadBalancer, graphTargetGroup)

    const graphServiceResourceName = getResourceName('analytics-graph-svc')
    const graphService = new aws.ecs.Service(graphServiceResourceName, {
      cluster: cluster.arn,
      deploymentCircuitBreaker: {
        enable: true,
        rollback: true,
      },
      desiredCount: 1,
      deploymentMaximumPercent: 100,
      deploymentMinimumHealthyPercent: 0,
      enableEcsManagedTags: true,
      forceNewDeployment: true,
      healthCheckGracePeriodSeconds: 20,
      launchType: 'EC2',
      loadBalancers: [
        {
          containerName: graphTaskDefinition.family,
          containerPort: 8080,
          targetGroupArn: graphTargetGroup.arn,
        },
      ],
      name: graphServiceResourceName,
      taskDefinition: graphTaskDefinition.arn,
      tags: getTags(tags),
    })
    applyGraphServiceAutoscaling(config, graphService)

    return cluster 
}
