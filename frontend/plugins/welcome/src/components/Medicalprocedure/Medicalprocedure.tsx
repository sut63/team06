import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link as RouterLink } from 'react-router-dom';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
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

import { Theme, createStyles } from '@material-ui/core/styles';
import { ContentHeader, Link } from '@backstage/core';
import { Cookies } from '../../Cookie';


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



interface medicalProcedure {
    Patients: number;
    Proceduretypes: number;
    Doctors: number;
    Addedtime: Date;
}

const medicalProcedure: FC<{}> = () => {

    const classes = useStyles();
    const api = new DefaultApi();
    const cookie = new Cookies();

    var nursename = cookie.GetCookie("nursename");
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [proceduretypes, setProceduretypes] = React.useState<EntProcedureType[]>([]);

    const [doctorID, setDoctorID] = React.useState(Number);
    const [patientID, setPatientID] = React.useState(Number);

    const [orders, setOrders] = React.useState(String);
    const [proceduretypeID, setProceduretypeID] = React.useState(Number);
    const [rooms, setRooms] = React.useState(String);
    const [descriptions, setDescriptions] = React.useState(String);

    const [orderError, setOrderError] = React.useState(String);
    const [roomError, setRoomError] = React.useState(String);
    const [descripeError, setDescripeError] = React.useState(String);

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
    const orderHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        checkpattern("order", event.target.value as string);
        setOrders(event.target.value as string);
    };
    const roomHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        checkpattern("room", event.target.value as string);
        setRooms(event.target.value as string);
    };
    const descripeHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        checkpattern("descripe", event.target.value as string);
        setDescriptions(event.target.value as string);

    };
    const checkOrder = (value: string) => {
        return value === "";
    }
    const checkRoom = (value: string) => {
        return value === "";
    }
    const checkDescripe = (value: string) => {
        return value === "";
    }

    const checkerr = (field: string) => {
        switch (field) {
            case 'procedureOrder':
                alertMessage("error", "รหัส order ขึ้นต้นด้วย UNS ตามด้วยตัวเลข 6 ตัว");
                return;
            case 'procedureRoom':
                alertMessage("error", "รหัส room เป็นเลข 4 ตัว");
                return;
            case 'procedureDescripe':
                alertMessage("error", "กรุณาใส่ Description");
                return;
            default:
                alertMessage("error", "Save Failed");
                return;
        }
    }

    const checkpattern = (id: string, value: string) => {
        switch (id) {
            case 'order':
                checkOrder(value) ? setOrderError('') : setOrderError('รหัส order ขึ้นต้นด้วย UNS ตามด้วยตัวเลข 6 ตัว');
                return;
            case 'room':
                checkRoom(value) ? setRoomError('') : setRoomError('รหัส room เป็นเลข 4 ตัว');
                return;
            case 'descripe':
                checkDescripe(value) ? setDescripeError('') : setDescripeError('โปรดใส่ข้อความ');
                return;
            default:
                return;
        }
    }


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


    function clear() {
        setOrders('');
        setDoctorID(0);
        setPatientID(0);
        setProceduretypeID(0);
        setRooms('');
        setDescriptions('');
    }
    function Create() {

        checkpattern("order", orderError);
        checkpattern("room", roomError);
        checkpattern("descripe", descripeError);

        const apiUrl = 'http://localhost:8080/api/v1/medicalprocedure';
        const medicalprocedure = {
            Orders: orders,
            Doctors: doctorID,
            Patients: patientID,
            Rooms: rooms,
            Proceduretypes: proceduretypeID,
            Descriptions: descriptions,
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
                    Toast.fire({
                        icon: 'success',
                        title: 'Save Successfully',
                    });
                    clear();
                } else {
                    checkerr(data.error.Name);
                }

            });
    }
    const alertMessage = (icon: any, title: any) => {
        Toast.fire({
            icon: icon,
            title: title,
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
                    <text> <strong>{nursename}</strong></text>
                    <br /><br />
                    <Button
                        style={{ backgroundColor: '#e57373	' }}
                        component={RouterLink}
                        to="/searchmedicalprocedure"
                        variant="contained"
                        startIcon={<SearchTwoToneIcon />}
                    >
                        search
              </Button>
              &nbsp;&nbsp;
                    <Button
                        component={RouterLink}
                        to="/loginmedicalprocedure"
                        variant="contained"

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
                            <TextField
                                error={orderError ? true : false}
                                id="orders"
                                label="Order"
                                name="Order"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={orderError}
                                value={orders}
                                onChange={orderHandle}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                            />
                        </FormControl>

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
                            variant="filled"
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

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <TextField
                                error={roomError ? true : false}
                                id="rooms"
                                label="Room"
                                name="Room"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={roomError}
                                value={rooms}
                                onChange={roomHandle}
                            />
                        </FormControl>

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <TextField
                                error={descripeError ? true : false}
                                id="description"
                                label="Descriptions"
                                name="Order"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={descripeError}
                                value={descriptions}
                                onChange={descripeHandle}
                            />
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
                                    onClick={() => {
                                        Create()
                                    }}
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
