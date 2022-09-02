

export enum SockType {
    low,
    high,
    none
}

export class Sock {
    id: string = '';
    shoeSize: number = 0;
    description: string = ''; 
    color: string = '';
    picture: string = '';
    type: SockType = SockType.none;
    owner: string='';
}