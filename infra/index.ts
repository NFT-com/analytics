import * as console from 'console'
import fs from 'fs'
import * as process from 'process'
import * as upath from 'upath'
import * as pulumi from '@pulumi/pulumi'
import { SharedInfraOutput, sharedOutputFileName } from './defs'
import { createSharedInfra } from './shared'
import { createAnalyticEcsCluster } from './analytics'

export const sharedOutToJSONFile = (outMap: pulumi.automation.OutputMap): void => {
  const analyticECRRepo = outMap.analyticECRRepo.value
  const publicSubnets = outMap.publicSubnetIds.value
  const vpcId = outMap.vpcId.value
  const webSGId = outMap.webSGId.value
  const sharedOutput: SharedInfraOutput = {
    analyticECRRepo,
    publicSubnets,
    vpcId,
    webSGId,
  }
  const file = upath.joinSafe(__dirname, sharedOutputFileName)
  fs.writeFileSync(file, JSON.stringify(sharedOutput))
}

const main = async (): Promise<any> => {
  const args = process.argv.slice(2)
  const deployShared = args?.[0] === 'deploy:shared' || false
  const deployAnalytics = args?.[0] === 'deploy:analytics' || false

  if (deployShared) {
    return createSharedInfra(true)
      .then(sharedOutToJSONFile)
  }
  
  if (deployAnalytics) {
    return createAnalyticEcsCluster()
  }
}

main()
  .catch((err) => {
    console.error(err)
    process.exit(1)
  })

