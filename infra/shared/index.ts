import * as upath from 'upath'
import { deployInfra, pulumiOutToValue, getResourceName } from '../helper'
import { createRepositories } from './ecr'
import { createSecurityGroups } from './security-group'
import * as pulumi from "@pulumi/pulumi";

const pulumiProgram = async (): Promise<Record<string, any> | void> => {
  const config = new pulumi.Config()
  const org = 'nftcom'
  const project = pulumi.getProject() //nftcom
  const indexerStack = getResourceName('indexer.shared.us-east-1')
  const indexerStackRef = new pulumi.StackReference(indexerStack)
  
  const zones = config.require('availabilityZones').split(',')
  const vpc = await pulumiOutToValue(indexerStackRef.getOutput('vpcId'))
  const subnets = indexerStackRef.getOutput('publicSubnetIds')
  const sgs = createSecurityGroups(config, vpc)
  const dbJob = indexerStackRef.getOutput('jobDbHost')
  const dbEvent = indexerStackRef.getOutput('eventDbHost')
  const dbGraph = indexerStackRef.getOutput('graphDbHost')
  const { analytics } = createRepositories()

  return {
    jobDbHost: dbJob,
    eventDbHost: dbEvent,
    graphDbHost: dbGraph,
    analyticECRRepo: analytics.name,
    publicSubnetIds: subnets,
    vpcId: vpc,
    webSGId: sgs.web.id,
  }
}

export const createSharedInfra = (
  preview?: boolean,
): Promise<pulumi.automation.OutputMap> => {
  const stackName = `${process.env.STAGE}.analytics.shared.${process.env.AWS_REGION}`
  const workDir = upath.joinSafe(__dirname, 'stack')
  return deployInfra(stackName, workDir, pulumiProgram, preview)
}
