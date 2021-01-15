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
    EntAdminrepair,
    EntAdminrepairFromJSON,
    EntAdminrepairFromJSONTyped,
    EntAdminrepairToJSON,
    EntBrand,
    EntBrandFromJSON,
    EntBrandFromJSONTyped,
    EntBrandToJSON,
    EntPersonal,
    EntPersonalFromJSON,
    EntPersonalFromJSONTyped,
    EntPersonalToJSON,
    EntTypeproduct,
    EntTypeproductFromJSON,
    EntTypeproductFromJSONTyped,
    EntTypeproductToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProductEdges
 */
export interface EntProductEdges {
    /**
     * 
     * @type {EntBrand}
     * @memberof EntProductEdges
     */
    brand?: EntBrand;
    /**
     * 
     * @type {EntPersonal}
     * @memberof EntProductEdges
     */
    personal?: EntPersonal;
    /**
     * Product holds the value of the product edge.
     * @type {Array<EntAdminrepair>}
     * @memberof EntProductEdges
     */
    product?: Array<EntAdminrepair>;
    /**
     * 
     * @type {EntTypeproduct}
     * @memberof EntProductEdges
     */
    typeproduct?: EntTypeproduct;
}

export function EntProductEdgesFromJSON(json: any): EntProductEdges {
    return EntProductEdgesFromJSONTyped(json, false);
}

export function EntProductEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProductEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'brand': !exists(json, 'brand') ? undefined : EntBrandFromJSON(json['brand']),
        'personal': !exists(json, 'personal') ? undefined : EntPersonalFromJSON(json['personal']),
        'product': !exists(json, 'product') ? undefined : ((json['product'] as Array<any>).map(EntAdminrepairFromJSON)),
        'typeproduct': !exists(json, 'typeproduct') ? undefined : EntTypeproductFromJSON(json['typeproduct']),
    };
}

export function EntProductEdgesToJSON(value?: EntProductEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'brand': EntBrandToJSON(value.brand),
        'personal': EntPersonalToJSON(value.personal),
        'product': value.product === undefined ? undefined : ((value.product as Array<any>).map(EntAdminrepairToJSON)),
        'typeproduct': EntTypeproductToJSON(value.typeproduct),
    };
}


