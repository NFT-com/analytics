import * as upath from 'upath'
import { deployInfra, pulumiOutToValue, getStage } from '../helper'
import { createRepositories } from './ecr'
import { createSecurityGroups } from './security-group'
import * as pulumi from '@pulumi/pulumi';

const pulumiProgram = async (): Promise<Record<string, any> | void> => {
  const config = new pulumi.Config()
  const stage = getStage()
  const sharedStack = new pulumi.StackReference(`${stage}.indexer.shared.us-east-1`);
  const vpc = sharedStack.getOutput('vpcId') // 'vpc-068564e7eded7ab8b'
  const subnets =  sharedStack.getOutput('subnets') //  ['subnet-0e2f01ec6714dc53f','subnet-0c8aa8a71e35104fc','subnet-08ea44006fecc2ab2']
  console.log(`vpc is ${vpc}`)
  const vpcVal: pulumi.Output<string> = await pulumiOutToValue(vpc)
  const subnetVal: pulumi.Output<string> = await pulumiOutToValue(subnets)
  console.log(`vpc.tostring  is ${vpc.toString()}`)
  console.log(`vpcVal (pulumi.output<string>) is ${vpcVal}`)
  console.log(`vpcVal.tostring  is ${vpcVal.toString()}`)


  const sgs = createSecurityGroups(config, vpcVal.toString()) //hardcode test
  const { analytics } = createRepositories()

  return {
    analyticECRRepo: analytics.name,
    publicSubnetIds: subnetVal,
    vpcId: vpcVal,
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
