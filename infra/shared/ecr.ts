import * as aws from '@pulumi/aws'

import { getResourceName } from '../helper'

export type RepositoryOut = {
  analytics: aws.ecr.Repository
}

export const createAnalyticsRepository = (): aws.ecr.Repository => {
  return new aws.ecr.Repository('ecr_analytics', {
    name: getResourceName('analytics'),
    imageScanningConfiguration: {
      scanOnPush: true,
    },
  })
}

export const createRepositories = (): RepositoryOut => {
  const analyticsRepo = createAnalyticsRepository()
  return {
    analytics: analyticsRepo,
  }
}
