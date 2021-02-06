import React, { useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, ContentHeader, pageTheme, } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, TableCell, Switch, } from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import moment from 'moment';
import Swal from 'sweetalert2';
import { Cookies } from '../../Cookie';

import { EntPatient } from '../../api/models/EntPatient';
import { EntDepartment } from '../../api/models/EntDepartment';
import { EntUrgencyLevel } from '../../api/models/EntUrgencyLevel';

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
            marginTop: theme.spacing(3),
        },
        textField: {
            flexWrap: 'wrap',
            margin: theme.spacing(3),
        },
        text: {
            color: '#ffffff',
            fontSize: 18,
        }
    }),
);

export default function Create() {

    const classes = useStyles();
    const api = new DefaultApi();
    const cookie = new Cookies();
    const [loading, setLoading] = useState(true);

    var nursename = cookie.GetCookie("nursename");
    var nurseID = cookie.GetCookie("nurseID");

    //query
    const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [urgencyLevels, setUrgencyLevels] = React.useState<EntUrgencyLevel[]>([]);

    //input
    const [departmentID, setDepartmentID] = React.useState(Number);
    const [patientID, setPatientID] = React.useState(Number);
    const [urgencyLevelID, setUrgencyLevelID] = React.useState(Number);
    const [symptom, setSymptom] = React.useState(String);
    const [height, setHeight] = React.useState(String);
    const [weight, setWeight] = React.useState(String);
    const [pressure, setPressure] = React.useState(String);
    const [triageDate, settriageDate] = useState('');

    //error
    const [symptomError, setSymptomError] = React.useState(String);
    const [heightError, setHeightError] = React.useState(String);
    const [weightError, setWeightError] = React.useState(String);
    const [pressureError, setPressureError] = React.useState(String);

    //Aleart
    const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: toast => {
            toast.addEventListener('mouseenter', Swal.stopTimer);
            toast.addEventListener('mouseleave', Swal.resumeTimer);
        },
    });

    const aleartMessage = (icon: any, title: any) => {
        Toast.fire({
            icon: icon,
            title: title,
        });
    }


    //getdataToCombobox
    useEffect(() => {
        const getDepartments = async () => {
            const res = await api.listDepartment({ limit: 1000, offset: 0 });
            setDepartments(res);
        };

        const getPatients = async () => {
            const res = await api.listPatient({ limit: 1000, offset: 0 });
            setPatients(res);
        };

        const getUrgencyLevels = async () => {
            const res = await api.listUrgencylevel({ limit: 1000, offset: 0 });
            setUrgencyLevels(res);
        };

        var date = moment().format();

        settriageDate(date);
        getDepartments();
        getPatients();
        getUrgencyLevels();
        setLoading(false);
        if (nursename == "") {
            Swal.fire({
                title: 'เข้าสู่ระบบก่อนใช้งาน',
                position: 'center',
                showConfirmButton: false,
                timer: 3000,
                timerProgressBar: true,
                didOpen: toast => {
                    toast.addEventListener('mouseenter', Swal.stopTimer);
                    toast.addEventListener('mouseleave', Swal.resumeTimer);
                },
            });
            setTimeout(() => {
                window.location.replace("http://localhost:3000/triageresultlogin");
            }, 3000);
        }
    }, [loading]);

    //handle
    const departmentHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setDepartmentID(event.target.value as number);
    };
    const patientHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPatientID(event.target.value as number);
    };
    const urgencyLevelHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setUrgencyLevelID(event.target.value as number);
    };
    const symptomHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setSymptom(event.target.value as string);
        CheckPattern("symptom", event.target.value as string);
    };
    const heightHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setHeight(event.target.value as string);
        CheckPattern("height", event.target.value as string);
    };
    const weightHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setWeight(event.target.value as string);
        CheckPattern("weight", event.target.value as string);
    };
    const pressureHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPressure(event.target.value as string);
        CheckPattern("pressure", event.target.value as string);
    };

    //checkValue
    const checkSymptom = (value: string) => {
        return value.length > 0;
    }
    const checkHeight = (value: string) => {
        return value.length > 0 && Number(value) && Number(value) > 0;
    }
    const checkWeight = (value: string) => {
        return value.length > 0 && Number(value) && Number(value) > 0;
    }
    const checkPressure = (value: string) => {
        return value.length > 0 && Number(value) && Number(value) > 0;
    }

    //CheckPattern
    const CheckPattern = (id: string, value: string) => {
        switch (id) {
            case 'symptom':
                checkSymptom(value) ? setSymptomError('') : setSymptomError('กรุณากรอกอาการ!');
                return;
            case 'height':
                checkHeight(value) ? setHeightError('') : setHeightError('ผิดพลาด กรุณากรอกใหม่!');
                return;
            case 'weight':
                checkWeight(value) ? setWeightError('') : setWeightError('ผิดพลาด กรุณากรอกใหม่!');
                return;
            case 'pressure':
                checkPressure(value) ? setPressureError('') : setPressureError('ผิดพลาด กรุณากรอกใหม่!');
                return;
            default:
                return;
        }

    }

    //Clear Data in field
    function Clear() {
        setDepartmentID(0);
        setPatientID(0);
        setUrgencyLevelID(0);
        setHeight('');
        setWeight('');
        setPressure('');
        setSymptom('');
    }

    //logout
    const Logout = async () => {
        cookie.SetCookie('nursename', "", 30)
        console.log("Logout success");
        aleartMessage("success", "ออกจากระบบสำเร็จ!");
    }

    //create
    const Create = async () => {

        CheckPattern("symptom", symptom);
        CheckPattern("height", height);
        CheckPattern("weight", weight);
        CheckPattern("pressure", pressure);


        const apiUrl = 'http://localhost:8080/api/v1/triageresults';
        const triageresult = {
            symptom: symptom,
            height: height,
            weight: weight,
            pressure: pressure,
            TriageDate: triageDate,
            Patient: patientID,
            Nurse: Number(nurseID),
            Department: departmentID,
            UrgencyLevel: urgencyLevelID,
        };
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(triageresult),
        };

        console.log(triageresult);

        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status === true) {
                    Toast.fire({
                        icon: 'success',
                        title: 'บันทึกข้อมูลสำเร็จ',
                    });
                    Clear();
                } else {
                    checkCaseError(data.error.Name);
                }
            });
    }

    const checkCaseError = (field: string) => {
        switch (field) {
            case 'symptom':
                aleartMessage("error", "กรุณากรอกอาการ");
                return;
            case 'height':
                aleartMessage("error", "ค่าส่วนสูงที่กรอกไม่ถูกต้อง");
                return;
            case 'weight':
                aleartMessage("error", "ค่าน้ำหนักที่กรอกไม่ถูกต้อง");
                return;
            case 'pressure':
                aleartMessage("error", "ค่าความดันที่กรอกไม่ถูกต้อง");
                return;
            default:
                aleartMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
                return;
        }
    }

    return (
        <Page theme={pageTheme.home}>

            <Header
                title={`ระบบบันทึกผลการคัดกรอง`}
            >
                <TableCell align={
                    "center"} >
                    <text
                        className={classes.text}
                    >
                        {nursename}
                    </text>
                    <br /><br />

                    <br /><br />
                    <Button
                        onClick={
                            () => {
                                Logout();
                            }
                        }
                        component={RouterLink}
                        to="/triageresultlogin"
                        variant="contained"
                        color="secondary"
                    >
                        Logout
                </Button>
                </TableCell>
            </Header>

            <Content>
                <div className={classes.root}>
                    <form noValidate autoComplete="off">
                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Patient</InputLabel>
                            <Select
                                id="patients"
                                label="Patient"
                                name="owner"
                                value={patientID}
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
                            <TextField
                                error={symptomError ? true : false}
                                id="symptom"
                                label="Symptom"
                                name="symptom"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={symptomError}
                                value={symptom}
                                onChange={symptomHandle}
                            />
                        </FormControl>

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <TextField
                                error={heightError ? true : false}
                                id="height"
                                label="Height"
                                name="height"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={heightError}
                                value={height}
                                onChange={heightHandle}
                            />
                        </FormControl>

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <TextField
                                error={weightError ? true : false}
                                id="weight"
                                label="Weight"
                                name="weight"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={weightError}
                                value={weight}
                                onChange={weightHandle}
                            />
                        </FormControl>

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <TextField
                                error={pressureError ? true : false}
                                id="pressure"
                                label="Pressure"
                                name="pressure"
                                variant="outlined"
                                type="string"
                                size="medium"
                                helperText={pressureError}
                                value={pressure}
                                onChange={pressureHandle}
                            />
                        </FormControl>

                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Department</InputLabel>
                            <Select
                                id="departments"
                                label="Department"
                                name="admit"
                                value={departmentID}
                                onChange={departmentHandle}
                            >
                                {departments.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.departmentName}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>

                        <FormControl
                            className={classes.margin}
                            fullWidth
                            variant="outlined"
                        >
                            <TextField
                                id="nurses"
                                label="Nurse"
                                name="recorder"
                                variant="outlined"
                                size="medium"
                                value={nursename}
                                disabled
                            />
                        </FormControl>

                        <FormControl
                            className={classes.margin}
                            fullWidth
                            variant="outlined"
                        >
                            <InputLabel>UrgencyLeval</InputLabel>
                            <Select
                                id="urgencyLevels"
                                label="UrencyLevel"
                                name="urgencyLevel"
                                value={urgencyLevelID}
                                onChange={urgencyLevelHandle}
                            >
                                {urgencyLevels.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.urgencyName}
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
                                id="datetime-local"
                                label="Date"
                                name="added"
                                variant="outlined"
                                size="medium"
                                value={moment(triageDate).format('LLLL')}
                                disabled
                            />
                        </FormControl>

                        <div className={classes.root}>
                            <TableCell>
                                <Button
                                    onClick={Create}
                                    variant="contained"
                                    size="large"
                                    color="primary"
                                >
                                    SAVE
                                </Button>
                            </TableCell>

                            <TableCell>
                                <Button
                                    component={RouterLink}
                                    to='/triageresultsearch'
                                    variant="contained"
                                    size="large"
                                    color="secondary"
                                >
                                    RESULT
                                </Button>
                            </TableCell>
                        </div>
                    </form>
                </div>
                <TableCell></TableCell>
            </Content>
        </Page>
    );
}
