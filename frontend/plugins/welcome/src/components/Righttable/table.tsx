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
import moment from 'moment';

import { Link as RouterLink } from 'react-router-dom';
import {Link,} from '@backstage/core';
import { EntRightToTreatment } from '../../api/models/EntRightToTreatment'; 


const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);

 const [righttotreatment, setRighttotreatment] = useState<EntRightToTreatment[]>([]);
 useEffect(() => {
  const getRighttotreatment = async () => {
    const le = await api.listRighttotreatment({ limit: 1000, offset:  0});
    setLoading(false);
    setRighttotreatment(le);
   
  };
  getRighttotreatment();
}, [loading]);

 return (

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">ผู้ป่วย</TableCell>
           <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
           <TableCell align="center">อายุ</TableCell>
           <TableCell align="center">ประเภทสิทธิการรักษา</TableCell>
           <TableCell align="center">โรงพยาบาลที่ใช้สิทธิ</TableCell>
           <TableCell align="center">เบอร์โทรศัพท์</TableCell>
           <TableCell align="center">วันที่เริ่มใช้สิทธิ</TableCell>
           <TableCell align="center">วันที่หมดสิทธิ</TableCell>
           <TableCell align="center">
            <Link component={RouterLink} to='/create'>
                                <Button 
                                    variant="contained"
                                    size="large"
                                    color="secondary"
                                    >
                                        เพิ่มบันทึก
                                </Button>
            </Link>
           </TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         
         {righttotreatment.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
             <TableCell align="center">{item.idennum}</TableCell>
             <TableCell align="center">{item.age}</TableCell>
             <TableCell align="center">{item.edges?.rightToTreatmentType?.typeName}</TableCell>
             <TableCell align="center">{item.edges?.hospital.hospitalName}</TableCell>
             <TableCell align="center">{item.tel}</TableCell>
             <TableCell align="center">{moment(item.startTime).format('DD/MM/YYYY')}</TableCell>
             <TableCell align="center">{moment(item.endTime).format('DD/MM/YYYY')}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}

