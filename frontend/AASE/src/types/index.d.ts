export type Flight = {
    id: number;
    origin: string;
    destination: string;
    aircraft_type: string;
    scheduled_departure_time: string;
    scheduled_arrival_time: string;
}

export type Seat = {
    id: number;
    flight_id: number;
    flight: Flight;
    account: Account;
    seat_number: string;
    assigned_pnr: string;
    for_sale: boolean;
    min_price: number;
    buy_now_price: number;
    current_bid: number;
    last_bidder: string;
    bid_time_end: string;
}

export type Account = {
    id: string;
    email: string;
    first_name: string;
    last_name: string;
    money: number;
}