import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { Link,} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import { EntAppointmentResults } from '../../api/models/EntAppointmentResults'; 
import moment from "moment";
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);

 const [appointmentresultss, setAppointmentResults] = useState<EntAppointmentResults[]>([]);

 useEffect(() => {
  const getAppointmentResults = async () => {
    const app = await api.listAppointmentresults({ limit: 1000, offset:  0});
    setLoading(false);
    setAppointmentResults(app);
  };
  getAppointmentResults();

}, [loading]);

 return (
   <TableContainer component={Paper}> 
     <Table className={classes.table} aria-label="simple table">
       <TableHead >
         <TableRow>
           <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">พยาบาลผู้บันทึก</TableCell>
           <TableCell align="center">ผู้ป่วย</TableCell>
           <TableCell align="center">แพทย์</TableCell>
           <TableCell align="center">ห้องตรวจ</TableCell>
           <TableCell align="center">วัน/เวลานัดหมาย</TableCell>
           <TableCell align="center">วัน/เวลาที่บันทึก</TableCell>
           <TableCell align="center">
            <Link component={RouterLink} to='/createappointment'>
                                <Button 
                                    variant="contained"
                                    size="large"
                                    color="secondary"
                                    >
                                        เพิ่มรายการนัดหมาย
                                </Button>
            </Link>
           </TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {appointmentresultss.sort().map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.appointmentResultsToNurse?.nurseName}</TableCell>
             <TableCell align="center">{item.edges?.appointmentResultsToPatient?.patientName}</TableCell>
             <TableCell align="center">{item.edges?.appointmentResultsToDoctor?.doctorName}</TableCell>
             <TableCell align="center">{item.edges?.appointmentResultsToRoom?.roomName}</TableCell>
             <TableCell align="center">{moment(item.addtimeAppoint).format("LLLL")}</TableCell>
             <TableCell align="center">{moment(item.addtimeSave).format("LLLL")}</TableCell>
           </TableRow>         
          ))}          
       </TableBody>
     </Table>
   </TableContainer>
 );
}

