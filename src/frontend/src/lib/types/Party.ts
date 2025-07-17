import type PartyTypes from "$lib/enums/PartyTypes.js"

export default interface Party {
    name: string
    type: PartyTypes
}