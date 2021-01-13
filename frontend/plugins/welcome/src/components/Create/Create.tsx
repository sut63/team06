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

import { EntHospital } from '../../api/models/EntHospital'; 
import { EntPatient } from '../../api/models/EntPatient'; 
import { EntRightToTreatmentType} from '../../api/models/EntRightToTreatmentType'; 
  

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
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
});
interface RightToTreatment {
    hospital: number;
    patient: number;
    rightToTreatmentType: number;
    Starttime: Date;
    endtime: Date;
}

const RightToTreatment: FC<{}> = () => {

    const classes = useStyles();
    const api = new DefaultApi();
    const [right, setRightToTreatment] = React.useState<
        Partial<RightToTreatment>
    >({});
   

    const [hospitals, setHospitals] = useState<EntHospital[]>([]);
    const [patients, setPatients] = useState<EntPatient[]>([]);
    const [righttotreatmentypes, setRightToTreatmentTypes] = useState<EntRightToTreatmentType[]>([]);

    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);
    
    const [hospitalid, sethospital] = useState(Number);
    const [patientid, setpatient] = useState(Number);
    const [righttotreatmentypeid, setrighttotreatmentype] = useState(Number);

    const [StartTime, setstarttime] = useState('');
    const [EndTime, setendtime] = useState('');

      useEffect(() => {
        var date = moment().format();
        setstarttime(StartTime);

        var date = moment().format();
        setendtime(EndTime);

      const getHospital = async () => {
          const d = await api.listHospital({ limit: 10, offset: 0 });
          setLoading(false)
          setHospitals(d);
        };
        getHospital();
  
        const getPatient = async () => {
          const p = await api.listPatient({ limit: 100, offset: 0 });
          setLoading(false)
          setPatients(p);
        };
        getPatient();
  
        const getRighttotreatmenttype = async () => {
          const t = await api.listRighttotreatmenttype({ limit: 10, offset: 0 });
          setLoading(false)
          setRightToTreatmentTypes(t);
      };
      getRighttotreatmenttype();
  
  }, [loading]);

  const HospitalhandleChange=(event: React.ChangeEvent<{ value: unknown; }>) => {
    sethospital(event.target.value as number);
  };
  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpatient(event.target.value as number);
  };
  const RightToTreatmentTypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setrighttotreatmentype(event.target.value as number);
  };
  const StarttimehandleChange = (event: any) => {
      setstarttime(event.target.value as string);
    };
    const EndtimeThandleChange = (event: any) => {
      setendtime(event.target.value as string);
    };


    function Create() {
        const apiUrl = 'http://localhost:8080/api/v1/righttotreatments';
        const righttotreatment = {
        hospital: hospitalid,
        patient: patientid,
        rightToTreatmentType: righttotreatmentypeid,
        starttime : StartTime + "T00:00:00+00:00",
        endtime : EndTime + "T00:00:00+00:00",
        };
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(righttotreatment),
        };
        console.log(righttotreatment);


        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status === true) {
                    
                    Toast.fire({
                        icon: 'success',
                        title: 'บันทึกข้อมูลสำเร็จ',
                    });
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
                title={`ระบบบันทึกสิทธิการรักษา`}
                subtitle="ยินดีต้อนรับเข้าสู่ระบบระบบบันทึกสิทธิการรักษา"
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
                <ContentHeader title="ข้อมูลสิทธิการรักษา" >
                {status ? (
                    <div className={classes.root}>
                        {alert ? (
                            <Alert severity="success">
                                success  
                            </Alert>
                        ) : (
                            <Alert severity="warning" style={{ marginTop: 20 }}>
                               warning 
                            </Alert>
                        )}
                    </div>
                ) : null}
                </ContentHeader>
  
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
                                value={patientid}
                                onChange={PatienthandleChange}
                            >
                                {patients.map(item => (
                                    <MenuItem key={item.id} value={item.id}>
                                        {item.patientName}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
  
                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Right to treatment type</InputLabel>
                            <Select
                                id="righttotreatmentypes"
                                label="RightToTreatmenType"
                                value={righttotreatmentypeid}
                                onChange={RightToTreatmentTypehandleChange}
                            >
                                {righttotreatmentypes.map(item => (
                                    <MenuItem value={item.id}>
                                        {item.typeName}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
  
                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>Hospital</InputLabel>
                            <Select
                                id="hospitals"
                                label="Hospital"
                                value={hospitalid}
                                onChange={HospitalhandleChange}
                            >
                                {hospitals.map(item => (
                                    <MenuItem key={item.id} value={item.id}>
                                        {item.hospitalName}
                                    </MenuItem>
                                ))}
                            </Select>
                        </FormControl>
  
                        <b></b>
  
  
                      <TextField
                          className={classes.margin}
                          variant="outlined"
                          id="date"
                          label="Start Date"
                          type="date"
                          onChange={StarttimehandleChange}
                          
                          InputLabelProps={{
                          shrink: true,
                          }}
                          />
  
                      <TextField
                          className={classes.margin}
                          variant="outlined"
                          id="date"
                          label="End Date"
                          type="date"
                          onChange={EndtimeThandleChange}
                          
                          InputLabelProps={{
                          shrink: true,
                          }}
                          />
  
                  
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
                                <Link component={RouterLink} to='/Righttable'>
                                    <Button
                                        variant="contained"
                                        size="large"
                                        color="secondary"
                                    >
                                       Data result
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

  
export default RightToTreatment;
