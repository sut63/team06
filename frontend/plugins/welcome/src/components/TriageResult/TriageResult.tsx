import React, {useEffect, useState} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, ContentHeader, pageTheme,} from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, TableCell,} from '@material-ui/core';
//import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import moment from 'moment';
import Swal from 'sweetalert2';

import { EntPatient } from '../../api/models/EntPatient';
import { EntNurse } from '../../api/models/EntNurse';
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
    }),
);

export  default  function Create() {

    const classes = useStyles();
    const api = new DefaultApi();

    const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
    const [nurses, setNurses] = React.useState<EntNurse[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [urgencyLevels, setUrgencyLevels] = React.useState<EntUrgencyLevel[]>([]);

    const [departmentID, setDepartmentID] = React.useState(Number);
    const [nurseID, setNurseID] = React.useState(Number);
    const [patientID, setPatientID] = React.useState(Number);
    const [urgencyLevelID, setUrgencyLevelID] = React.useState(Number);
    const [symptom, setSymptom] = React.useState(String);
    const [triageDate, settriageDate] = useState('');

    //const [status, setStatus] = useState(false);
    //const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);

    //getdataToCombobox
    useEffect(() => {
        const getDepartments = async () => {
            const res = await api.listDepartment({ limit: 1000, offset:  0});
            setDepartments(res);
        };
    
        const getNurses = async () => {
            const res = await api.listNurse({ limit: 1000, offset:  0});
            setNurses(res);
        };
    
        const getPatients = async () => {
            const res = await api.listPatient({ limit: 1000, offset:  0});
            setPatients(res);
        };

        const getUrgencyLevels = async () => {
            const res = await api.listUrgencylevel({ limit: 1000, offset:  0});
            setUrgencyLevels(res);
        };

        var date = moment().format();

        settriageDate(date);
        getDepartments();
        getNurses();
        getPatients();
        getUrgencyLevels();
        setLoading(false);
    }, [loading]);

    //handle
    const departmentHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setDepartmentID(event.target.value as number);
      };
    const patientHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPatientID(event.target.value as number);
      };
    const nurseHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setNurseID(event.target.value as number);
      };
    const urgencyLevelHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setUrgencyLevelID(event.target.value as number);
      };
    const symptomHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setSymptom(event.target.value as string);
      };

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

    //Clear Data in field
    function Clear() {
        setDepartmentID(0);
        setPatientID(0);
        setNurseID(0);
        setUrgencyLevelID(0);
        setSymptom('');
    }
    //create
    const Create = async () => {
        const apiUrl = 'http://localhost:8080/api/v1/triageresults';
        const triageresult = {
        symptom :       symptom,
        TriageDate :    triageDate,
        Patient :       patientID,
        Nurse :         nurseID ,
        Department :    departmentID,
        UrgencyLevel :  urgencyLevelID,
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
                  Toast.fire({
                    icon: 'error',
                    title: 'บันทึกข้อมูลไม่สำเร็จ',
                });
              }
            });
    }
    
return (
        <Page theme={pageTheme.home}>
            
            <Header
                title={`ระบบบันทึกผลการคัดกรอง`}
            >
            <TableCell align={
                "center"} >

                <br/><br/>
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
                                id="symptom"
                                label="Symptom"
                                name="symptom"
                                variant="outlined"
                                type="string"
                                size="medium"
                                value={symptom}
                                onChange={symptomHandle}
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
                            <InputLabel>Nurse</InputLabel>
                            <Select
                                id="nurses"
                                label="Nurse"
                                name="recorder"
                                value={nurseID}
                                onChange={nurseHandle}
                            >
                                {nurses.map(item => {
                                    return (
                                    <MenuItem key={item.id} value={item.id}>
                                        {item.nurseName}
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
                                value={triageDate}
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
                                    to='/result'
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
