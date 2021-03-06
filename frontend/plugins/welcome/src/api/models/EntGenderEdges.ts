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
import {
    EntCustomer,
    EntCustomerFromJSON,
    EntCustomerFromJSONTyped,
    EntCustomerToJSON,
    EntPersonal,
    EntPersonalFromJSON,
    EntPersonalFromJSONTyped,
    EntPersonalToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntGenderEdges
 */
export interface EntGenderEdges {
    /**
     * Customer holds the value of the customer edge.
     * @type {Array<EntCustomer>}
     * @memberof EntGenderEdges
     */
    customer?: Array<EntCustomer>;
    /**
     * Personal holds the value of the personal edge.
     * @type {Array<EntPersonal>}
     * @memberof EntGenderEdges
     */
    personal?: Array<EntPersonal>;
}

export function EntGenderEdgesFromJSON(json: any): EntGenderEdges {
    return EntGenderEdgesFromJSONTyped(json, false);
}

export function EntGenderEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntGenderEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customer': !exists(json, 'customer') ? undefined : ((json['customer'] as Array<any>).map(EntCustomerFromJSON)),
        'personal': !exists(json, 'personal') ? undefined : ((json['personal'] as Array<any>).map(EntPersonalFromJSON)),
    };
}

export function EntGenderEdgesToJSON(value?: EntGenderEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customer': value.customer === undefined ? undefined : ((value.customer as Array<any>).map(EntCustomerToJSON)),
        'personal': value.personal === undefined ? undefined : ((value.personal as Array<any>).map(EntPersonalToJSON)),
    };
}


