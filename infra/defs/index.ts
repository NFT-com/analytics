export const sharedOutputFileName = 'shared-out.json'

export type SharedInfraOutput = {
  jobDbHost: string
  eventDbHost: string
  graphDbHost: string
  analyticECRRepo: string
  publicSubnets: string[]
  vpcId: string
  webSGId: string
}
