import React, {FC,useEffect, useState} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, ContentHeader, pageTheme,Link} from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, TableCell} from '@material-ui/core';
import { Alert ,AlertTitle } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import moment from 'moment'
import { EntPatient } from '../../api/models/EntPatient';
import { EntNurse } from '../../api/models/EntNurse';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntRoom } from '../../api/models/EntRoom';
import SaveIcon from '@material-ui/icons/Save';

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
            '& > span': {margin: theme.spacing(3),},
            
        },
        textField: {
            flexWrap: 'wrap',
            margin: theme.spacing(3),
        },
    }),
);

const CreateAppointmentResults: FC<{}> = () => {
  
    const classes = useStyles();
    const api = new DefaultApi();
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [nurses, setNurses] = React.useState<EntNurse[]>([]);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);
    const [rooms, setRooms] = React.useState<EntRoom[]>([]); 
    const [doctorID, setDoctorID] = React.useState(Number);
    const [nurseID, setNurseID] = React.useState(Number);
    const [patientID, setPatientID] = React.useState(Number);
    const [roomID, setRoomID] = React.useState(Number);
    const [addtimeappoint, setAddtimeAppoint] = useState(Date);
    const [addtimesave, setAddtimeSave] = useState('');
    
    const [success, setSuccess] = React.useState(Number);
    //get data to ui materials
    useEffect(() => {   
        var addtimesave = moment().format();
        const getNurses = async () => {
            const res = await api.listNurse({ limit: 1000, offset:  0});
            setNurses(res);
        };
    
        const getPatients = async () => {
            const res = await api.listPatient({ limit: 1000, offset:  0});
            setPatients(res);
        };
        const getDoctors = async () => {
            const res = await api.listDoctor({ limit: 1000, offset:  0});
            setDoctors(res);
        };
        const getRooms = async () => {
            const res = await api.listRoom({ limit: 1000, offset:  0});
            setRooms(res);
        };
        getNurses();
        getPatients();
        getDoctors();
        getRooms();
        setAddtimeSave(addtimesave);
    }, []);

  //handle
    
    const patientHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPatientID(event.target.value as number);
    };
    const nurseHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setNurseID(event.target.value as number);
    };
    const roomHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setRoomID(event.target.value as number);
    };   
    const doctorHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setDoctorID(event.target.value as number);
    };   
    const addtimesaveHandleChange = (event: any) => {
        setAddtimeSave(event.target.value as string);
    };
    const addtimeappointHandleChange = (event: any) => {
        setAddtimeAppoint(event.target.value as string);
    };
    
    
   function Create() {
    const apiUrl = 'http://localhost:8080/api/v1/appointmentresultss';
    const appointmentresults = {
        nurse :          nurseID,
        patient :        patientID,  
        doctor :         doctorID, 
        room :           roomID,
        addtimesave :    addtimesave,
        addtimeappoint : addtimeappoint + ":00+07:00"
    };

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(appointmentresults),
    };

    console.log(appointmentresults);

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status == true) {
            setSuccess(1)         
        } 
        else{
            setSuccess(2)
        }
      });
  }

    return ( 
        <Page theme={pageTheme.home}>
            
            <Header
                title={`ระบบนัดหมาย`}
                subtitle={`Appointment System`}
            >
            <TableCell>   
                <Link component={RouterLink} to='/'>
                    <Button 
                        variant="contained"
                        size="large"
                     >
                            -LOGOUT-
                    </Button>
                </Link>       
            </TableCell>
            </Header>

            <Content>
                <ContentHeader title="สร้างรายการนัดหมาย" >
                
                    <div className={classes.root}>
                        {success == 1 ? (
                            <Alert severity="success" > 
                                บันทึกข้อมูลสำเร็จ! — <strong>Success!</strong>
                            </Alert>
                        ) : <b></b>}
                    </div>
                    <div className={classes.root}>
                        {success == 2 ? (
                            <Alert severity="error" >
                                
                                    ERROR!---บันทึกข้อมูลไม่สำเร็จ! — <strong>check it out!</strong> 
                            </Alert>
                        ) : <b></b>}
                    </div>
                
                </ContentHeader>

                <div className={classes.root}>
                    <form noValidate autoComplete="off">                         
                    <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                            
                        >
                            <InputLabel>พยาบาลผู้ใช้ระบบ</InputLabel>
                            <Select
                                id="nurses"
                                label="Nurse"
                                name="nurse"
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
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel>เลือกผู้ป่วย</InputLabel>
                            <Select
                                id="patients"
                                label="Patient"
                                name="patient"
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
                            <InputLabel>เลือกแพทย์</InputLabel>
                            <Select
                                id="doctors"
                                label="Doctor"
                                name="doctor"
                                value={doctorID}
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
                            <InputLabel>เลือกห้องตรวจ</InputLabel>
                            <Select
                                id="rooms"
                                label="Room"
                                name="room"
                                value={roomID}
                                onChange={roomHandle}
                            >
                                {rooms.map(item => {
                                return (
                                <MenuItem key={item.id} value={item.id}>
                                    {item.roomName}
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
                                disabled
                                id="datetime-local"
                                label="วัน/เวลาที่บันทึก"
                                value={addtimesave}
                                onChange={addtimesaveHandleChange}
                                InputLabelProps={{
                                shrink: true
                                }}
                            />
                        </FormControl>

                        <FormControl 
                            fullWidth
                            className={classes.margin}
                            variant="outlined"                       
                            >
                            <TextField
                                id="datetime-local"
                                label="วัน/เวลานัดหมาย"
                                type="datetime-local" 
                                value={addtimeappoint}
                                onChange={addtimeappointHandleChange}
                                InputLabelProps={{
                                    shrink: true
                                }}
                            />
                        </FormControl>

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
                                    บันทึก
                                </Button>
                            </TableCell>

                            <TableCell>
                            <Link component={RouterLink} to='/appointmentresults'>
                                <Button 
                                    variant="outlined" 
                                    size="large"
                                    color="secondary"
                                    >
                                        ดูรายการนัดหมาย
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

export default CreateAppointmentResults;
