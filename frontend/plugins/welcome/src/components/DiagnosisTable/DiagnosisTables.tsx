import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatient } from '../../api/models/EntPatient';
import { EntTreatmentType } from '../../api/models/EntTreatmentType'; 
import { DefaultApi } from '../../api/apis';
import { EntDiagnosis } from '../../api/models/EntDiagnosis';



const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const api = new DefaultApi();

    const [diagnossis, setDiagnosis] = React.useState<EntDiagnosis[]>([]);
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [treatmenttype, SetType] = React.useState<EntTreatmentType[]>([]);
    
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getDoctors = async () => {
            const res = await api.listDoctor({ limit: 1000, offset: 0 });
            setDoctors(res);
        };
        const getPatients = async () => {
            const res = await api.listPatient({ limit: 1000, offset: 0 });
            setPatients(res);
        };
        const getTreatmentType = async () => {
            const res = await api.listTreatmentType({ limit: 1000, offset: 0 });
            SetType(res);
        };
        



        getDoctors();
        getPatients();
        getTreatmentType();
        
        
        
        setLoading(false);
    }, [loading]);

    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="center">ชื่อแพทย์</TableCell>
                        <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                        <TableCell align="center">อาการ</TableCell>
                        <TableCell align="center">ความคิดเห็นจากแพทย์</TableCell>
                        <TableCell align="center">วันที่</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {diagnossis.sort().map((item: any) => (
                        <TableRow key={item.id}>
                            
                            <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                            <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                            <TableCell align="center">{item.edges?.procedureType?.procedureName}</TableCell>
                            <TableCell align="center">{moment(item.addtime).format('LLLL')}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}
