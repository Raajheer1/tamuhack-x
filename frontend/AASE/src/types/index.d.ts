export type Flight = {
    id: number;
    origin: string;
    destination: string;
    aircraft_type: string;
    schedule_departure_time: string;
    schedule_arrival_time: string;
}

export type Seat = {
    id: number;
    flight_id: number;
    flight: Flight;
    seat_number: string;
    assigned_pnr: string;
    for_sale: boolean;
    min_price: number;
    buy_now_price: number;
    current_bid: number;
    last_bidder: string;
    bid_time_end: string;
}