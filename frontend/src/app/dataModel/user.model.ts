import { Address } from "./address.model";
export interface User {
    fisrtname: string;
    surname: string;
    username: string;
    shippingAddress: Address;
}