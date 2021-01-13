import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link as RouterLink } from 'react-router-dom';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  TableCell,
} from '@material-ui/core';
import Icon from '@material-ui/core/Icon';

import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import moment from 'moment';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatient } from '../../api/models/EntPatient';
import { EntProcedureType } from '../../api/models/EntProcedureType'; 
import { EntNurse } from '../../api/models/EntNurse'; 

import {  Theme, createStyles } from '@material-ui/core/styles';
import {  ContentHeader, Link } from '@backstage/core';


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',

        },
        margin: {
            display: 'flex',
            flexWrap: 'wrap',
            width: '70ch',
            margin: theme.spacing(3),
        },
        withoutLabel: {
            '& > span': { margin: theme.spacing(3), },

        },
        textField: {
            flexWrap: 'wrap',
            margin: theme.spacing(3),
        },
    }),
);
const Toast = Swal.mixin({
    toast: true,
    position: 'top',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
});
interface medicalProcedure {
    Patients: number;
    Proceduretypes: number;
    Doctors: number;
    Addedtime: Date;
}

const medicalProcedure: FC<{}> = () => {

    const classes = useStyles();
    const api = new DefaultApi();
    const [medical_pro, setMedicals] = React.useState<
        Partial<medicalProcedure>
    >({});
   
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [proceduretypes, setProceduretypes] = React.useState<EntProcedureType[]>([]);

    const [doctorID, setDoctorID] = React.useState(Number);
    const [patientID, setPatientID] = React.useState(Number);
    const [proceduretypeID, setProceduretypeID] = React.useState(Number);


    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);
    const [addtime, setCurrentDate] = useState('');

    useEffect(() => {
        var date = moment().format();
        const getDoctor = async () => {
            const res = await api.listDoctor({ limit: 1000, offset: 0 });
            setLoading(false);
            setDoctors(res);
        };
        getDoctor();

        const getPatient = async () => {
            const res = await api.listPatient({ limit: 1000, offset: 0 });
            setLoading(false);
            setPatients(res);
        };
        getPatient();

        const getPeocedure = async () => {
            const res = await api.listProceduretype({ limit: 1000, offset: 0 });
            setLoading(false)
            setProceduretypes(res);
        };
        getPeocedure();
        setCurrentDate(date);

    }, [loading]);

    const doctorHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setDoctorID(event.target.value as number);
    };
    const patientHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPatientID(event.target.value as number);
    };
    const proceduretypeHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setProceduretypeID(event.target.value as number);
    };
    

    function clear() {
        setMedicals({});
    }


    function Create() {
        const apiUrl = 'http://localhost:8080/api/v1/medicalprocedure';
        const medicalprocedure = {
            Doctors: doctorID,
            Patients: patientID,
            Proceduretypes: proceduretypeID,
            Addedtime: addtime,
        };
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(medicalprocedure),
        };
        console.log(medicalprocedure);


        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status === true) {
                    clear();
                    Toast.fire({
                        icon: 'success',
                        title: 'Save successfully',
                    });
                } else {
                    Toast.fire({
                        icon: 'error',
                        title: 'Save Failed',
                    });
                }
            });
    }

    return (
        <Page theme={pageTheme.home}>

            <Header
                title={`Medical Procedure`}
                subtitle="ระบบการบันทึกการทำหัตถการ"
            >
                <TableCell align={
                    "center"} >
                    <br /><br />
                    <Button
                        component={RouterLink}
                        to="/"
                        variant="contained"
                        color="secondary"
                    >
                        Logout
              </Button>
                </TableCell>
            </Header>

            <Content>
                <ContentHeader title=""> </ContentHeader>

                <div className={classes.root}>
                    <form noValidate autoComplete="off">
                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Doctor</InputLabel>
                            <Select
                                id="doctor"
                                label="Doctor"
                                name="doctor"
                                value={doctorID || ''}
                                onChange={doctorHandle}
                            >

                                {doctors.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.doctorName}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Patient</InputLabel>
                            <Select
                                id="patient"
                                label="Patient"
                                name="patients"
                                value={patientID || ''}
                                onChange={patientHandle}
                            >

                                {patients.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.patientName}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>


                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Procedure Type</InputLabel>
                            <Select
                                id="proceduretype"
                                label="Proceduretype"
                                name="proceduretype"
                                value={proceduretypeID || ''}
                                onChange={proceduretypeHandle}
                            >
                                {proceduretypes.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.procedureName}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>

                        <b></b>

                        <form noValidate>
                            <TextField disabled
                                className={classes.margin}
                                variant="outlined"
                                id="datetime-local"
                                label="Datetime"

                                value={addtime}


                            />
                        </form>

                        <div className={classes.root}>
                            <TableCell>
                                <Button
                                    onClick={Create}
                                    variant="contained"
                                    size="large"
                                    color="primary"
                                    startIcon={<SaveIcon />}
                                >
                                    SAVE
                                </Button>
                            </TableCell>

                            <TableCell>
                                <Link component={RouterLink} to='/medicalresults'>
                                    <Button
                                        variant="contained"
                                        size="large"
                                        color="secondary"
                                    >
                                        Result
                                </Button>
                                </Link>
                            </TableCell>
                            
                            

                        </div>

                    </form>


                </div>
                <TableCell></TableCell>
            </Content>
        </Page>
    );
}

export default medicalProcedure;
