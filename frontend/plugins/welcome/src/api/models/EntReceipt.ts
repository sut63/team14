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
    EntReceiptEdges,
    EntReceiptEdgesFromJSON,
    EntReceiptEdgesFromJSONTyped,
    EntReceiptEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntReceipt
 */
export interface EntReceipt {
    /**
     * AddedTime holds the value of the "added_time" field.
     * @type {string}
     * @memberof EntReceipt
     */
    addedTime?: string;
    /**
     * 
     * @type {EntReceiptEdges}
     * @memberof EntReceipt
     */
    edges?: EntReceiptEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntReceipt
     */
    id?: number;
}

export function EntReceiptFromJSON(json: any): EntReceipt {
    return EntReceiptFromJSONTyped(json, false);
}

export function EntReceiptFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntReceipt {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'added_time') ? undefined : json['added_time'],
        'edges': !exists(json, 'edges') ? undefined : EntReceiptEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntReceiptToJSON(value?: EntReceipt | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added_time': value.addedTime,
        'edges': EntReceiptEdgesToJSON(value.edges),
        'id': value.id,
    };
}


