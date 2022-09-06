

export enum SockType {
    low,
    medium,
    high,
    none
}

export function typeToString(sock: Sock): string {
    switch(sock.type) {
        case SockType.low:
            return 'Sockette';
        case SockType.medium:
            return 'Medium';
        case SockType.high:
            return 'Knee high';
        default:
            return 'none';
    }
}

export class Sock {
    id: string = '';
    shoeSize: number = 0;
    description: string = ''; 
    color: string = '';
    picture: string = '';
    type: SockType = SockType.none;
    owner: string='';
    acceptedList: string[] = new Array();
    refusedList: string[] = new Array();
    match: string = '';
    matchResult: string = '';
}

