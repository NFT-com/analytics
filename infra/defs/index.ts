export const sharedOutputFileName = 'shared-out.json'

export type SharedInfraOutput = {
  analyticECRRepo: string
  publicSubnets: string[]
  vpcId: string
  webSGId: string
}
