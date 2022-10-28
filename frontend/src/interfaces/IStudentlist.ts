import { ReportInterface } from "./IReport";
import { StatusInterface } from "./IStatus";

export interface StudentListInterface {
  ID?: number,
    
  SaveTime: Date | null,
  Reason: string | null,
  Amount: number | null,

  ReportID?: number,
  Report?: ReportInterface,  
  StatusID?: number,
  Status?: StatusInterface,
  AdminID?: number;
  }
  