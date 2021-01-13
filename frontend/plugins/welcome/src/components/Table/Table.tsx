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

import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatient } from '../../api/models/EntPatient';
import { EntProcedureType } from '../../api/models/EntProcedureType';
import { EntNurse } from '../../api/models/EntNurse'; 
import { EntMedicalProcedure } from '../../api/models/EntMedicalProcedure';


const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const api = new DefaultApi();

    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [proceduretypes, setProceduretypes] = React.useState<EntProcedureType[]>([]);
    const [medicalprocedures, setMedicalprocedures] = React.useState<EntMedicalProcedure[]>([]);
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
        const getProceduretypes = async () => {
            const res = await api.listProceduretype({ limit: 1000, offset: 0 });
            setProceduretypes(res);
        };
        const getMedicalprocedures = async () => {
            const res = await api.listMedicalprocedure({ limit: 50, offset: 0 });
            setLoading(false);
            setMedicalprocedures(res);
        };



        getDoctors();
        getPatients();
        getProceduretypes();
        getMedicalprocedures();
        setLoading(false);
    }, [loading]);

    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="center">No.</TableCell>
                        <TableCell align="center">Doctor</TableCell>
                        <TableCell align="center">Patient</TableCell>
                        <TableCell align="center">ProcedureType</TableCell>
                        <TableCell align="center">Datetime</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {medicalprocedures.sort().map((item: any) => (
                        <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
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
