export interface Pick {
  idPik: number;
  redBrPik: string; // Pick order
  brRunPik: string; // Pick round (can be first or second)
  godPik: string; // Pick year
  nedodListPik: boolean; // Is pick added to the untouchables list
  trgListPik: boolean; // Is pick added to the trade list
  idTim: number; // Team foreign key
}
