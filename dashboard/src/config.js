import * as development from './config/config.development'
import * as production from './config/config.production'

export const current = process.env.NODE_ENV === 'production' ? production.default : development.default