import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import moment from 'moment';

import { EntTriageResult } from '../../api/models/EntTriageResult';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
  
  const classes = useStyles();
  const api = new DefaultApi();

  const [triageresult,setTriageresult] = React.useState<EntTriageResult[]>([]);
  const [loading, setLoading] = useState(true);
 
  useEffect(() => {
    const getTriageResults = async () => {
      const res = await api.listTriageresult({ limit: 1000, offset:  0});
      setTriageresult(res);
    };

    getTriageResults();
    setLoading(false);
}, [loading]);



 return (
     <TableContainer component={Paper}>
       <Table className={classes.table} aria-label="simple table">
         <TableHead>
           <TableRow>
             <TableCell align="center">No.</TableCell>
             <TableCell align="center">Patient</TableCell>
             <TableCell align="center">Symptom</TableCell>
             <TableCell align="center">Height</TableCell>
             <TableCell align="center">Weight</TableCell>
             <TableCell align="center">Pressure</TableCell>
             <TableCell align="center">UrgencyLeval</TableCell>
             <TableCell align="center">Department</TableCell>
             <TableCell align="center">Nurse</TableCell>
             <TableCell align="center">Date</TableCell>
           </TableRow>
         </TableHead>
         <TableBody>
           {triageresult.sort().map((item: any) => (
             <TableRow key={item.id}>
               <TableCell align="center">{item.id}</TableCell>
               <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
               <TableCell align="center">{item.symptom}</TableCell>
               <TableCell align="center">{item.height}</TableCell>
               <TableCell align="center">{item.weight}</TableCell>
               <TableCell align="center">{item.pressure}</TableCell>
               <TableCell align="center">{item.edges?.urgencyLevel?.urgencyName}</TableCell>
               <TableCell align="center">{item.edges?.department?.departmentName}</TableCell>
               <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
               <TableCell align="center">{moment(item.triageDate).format('LLLL')}</TableCell>
             </TableRow>
           ))}
         </TableBody>
       </Table>
     </TableContainer>
 );
}
