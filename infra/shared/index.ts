import * as upath from 'upath'
import { deployInfra, getResourceName } from '../helper'
import { createRepositories } from './ecr'
import { createSecurityGroups } from './security-group'
import * as pulumi from '@pulumi/pulumi';

const pulumiProgram = async (): Promise<Record<string, any> | void> => {
  const config = new pulumi.Config()
  //const stackRefName = getResourceName('indexer.shared.us-east-1')
  const sharedStack = new pulumi.StackReference('dev.indexer.shared.us-east-1');
  const vpc = sharedStack.getOutput('vpcId').toString() // 'vpc-068564e7eded7ab8b'
  const subnets =  sharedStack.getOutput('subnets').toString() //  ['subnet-0e2f01ec6714dc53f','subnet-0c8aa8a71e35104fc','subnet-08ea44006fecc2ab2']

  const sgs = createSecurityGroups(config, vpc)
  const { analytics } = createRepositories()

  return {
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
