
import * as process from 'process'
import * as upath from 'upath'
import * as pulumi from '@pulumi/pulumi'

import { deployInfra, getSharedInfraOutput } from '../helper'
import { createEcsCluster, createEventsTaskDefinition, createAggregationTaskDefinition, createGraphTaskDefinition   } from './ecs'


const pulumiProgram = async (): Promise<Record<string, any> | void> => {
    const config = new pulumi.Config()
    const sharedInfraOutput = getSharedInfraOutput()
    const eventTaskDefinition = createEventsTaskDefinition(sharedInfraOutput)
    const aggregationTaskDefinition = createAggregationTaskDefinition(sharedInfraOutput)
    const graphTaskDefinition = createGraphTaskDefinition(sharedInfraOutput)
    createEcsCluster(config,sharedInfraOutput, eventTaskDefinition,aggregationTaskDefinition,graphTaskDefinition )
  }
  
  export const createAnalyticEcsCluster = (
    preview?: boolean,
  ): Promise<pulumi.automation.OutputMap> => {
    const stackName = `${process.env.STAGE}.analytics.${process.env.AWS_REGION}`
    const workDir = upath.joinSafe(__dirname, 'stack')
    return deployInfra(stackName, workDir, pulumiProgram, preview)
  }