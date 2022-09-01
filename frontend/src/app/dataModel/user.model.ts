import { Address } from "./address.model";
export class User {
    firstname: string = '';
    surname: string = '';
    username: string = '';
    password: string = '';
    address: Address = new Address();
}