interface Airport {
    code: string;
    city: string;
  }
  
  function getCityFromAirportCode(code: string): string | undefined {
    const airportCodes: Airport[] = [
      { code: "JFK", city: "New York" },
      { code: "LAX", city: "Los Angeles" },
      { code: "ORD", city: "Chicago" },
      { code: "ATL", city: "Atlanta" },
      { code: "DFW", city: "Dallas" },
      { code: "DEN", city: "Denver" },
      { code: "SFO", city: "San Francisco" },
      { code: "LAS", city: "Las Vegas" },
      { code: "SEA", city: "Seattle" },
      { code: "CLT", city: "Charlotte" },
      { code: "EWR", city: "Newark" },
      { code: "MCO", city: "Orlando" },
      { code: "PHX", city: "Phoenix" },
      { code: "IAH", city: "Houston" },
      { code: "MIA", city: "Miami" },
      { code: "BOS", city: "Boston" },
      { code: "MSP", city: "Minneapolis" },
      { code: "FLL", city: "Fort Lauderdale" },
      { code: "DTW", city: "Detroit" },
      { code: "PHL", city: "Philadelphia" },
      { code: "LGA", city: "New York" },
      { code: "BWI", city: "Baltimore" },
      { code: "SLC", city: "Salt Lake City" },
      { code: "SAN", city: "San Diego" },
      { code: "IAD", city: "Washington" },
      { code: "DCA", city: "Washington" },
      { code: "TPA", city: "Tampa" },
      { code: "PDX", city: "Portland" },
      { code: "HNL", city: "Honolulu" },
      { code: "STL", city: "St. Louis" },
      { code: "AUS", city: "Austin" },
      { code: "MDW", city: "Chicago" },
      { code: "DAL", city: "Dallas" },
      { code: "HOU", city: "Houston" },
      { code: "BNA", city: "Nashville" },
      { code: "OAK", city: "Oakland" },
      { code: "MSY", city: "New Orleans" },
      { code: "SMF", city: "Sacramento" },
      { code: "SJC", city: "San Jose" },
      { code: "SNA", city: "Santa Ana" },
      { code: "MCI", city: "Kansas City" },
      { code: "RSW", city: "Fort Myers" },
    ];
  
    const matchingAirport = airportCodes.find((airport) => airport.code === code);
    return matchingAirport ? matchingAirport.city : undefined;
  }
  
  export default getCityFromAirportCode;
  