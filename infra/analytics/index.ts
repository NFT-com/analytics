
import * as process from 'process'
import * as upath from 'upath'
import * as pulumi from '@pulumi/pulumi'

import { deployInfra, getSharedInfraOutput } from '../helper'
import { createEcsCluster  } from './ecs'


const pulumiProgram = async (): Promise<Record<string, any> | void> => {
    const config = new pulumi.Config()
    const sharedInfraOutput = getSharedInfraOutput()
    createEcsCluster(config,sharedInfraOutput)
  }
  
  export const createAnalyticEcsCluster = (
    preview?: boolean,
  ): Promise<pulumi.automation.OutputMap> => {
    const stackName = `${process.env.STAGE}.analytics.${process.env.AWS_REGION}`
    const workDir = upath.joinSafe(__dirname, 'stack')
    return deployInfra(stackName, workDir, pulumiProgram, preview)
  }