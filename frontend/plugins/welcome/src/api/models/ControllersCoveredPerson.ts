/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersCoveredPerson
 */
export interface ControllersCoveredPerson {
    /**
     * 
     * @type {number}
     * @memberof ControllersCoveredPerson
     */
    certificate?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersCoveredPerson
     */
    fund?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersCoveredPerson
     */
    patient?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersCoveredPerson
     */
    schemeType?: number;
}

export function ControllersCoveredPersonFromJSON(json: any): ControllersCoveredPerson {
    return ControllersCoveredPersonFromJSONTyped(json, false);
}

export function ControllersCoveredPersonFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersCoveredPerson {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'certificate': !exists(json, 'certificate') ? undefined : json['certificate'],
        'fund': !exists(json, 'fund') ? undefined : json['fund'],
        'patient': !exists(json, 'patient') ? undefined : json['patient'],
        'schemeType': !exists(json, 'schemeType') ? undefined : json['schemeType'],
    };
}

export function ControllersCoveredPersonToJSON(value?: ControllersCoveredPerson | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'certificate': value.certificate,
        'fund': value.fund,
        'patient': value.patient,
        'schemeType': value.schemeType,
    };
}


