import React, {FC,useEffect, useState} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, ContentHeader, pageTheme,Link} from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, TableCell, TableRow} from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import moment from 'moment'
import { EntPatient } from '../../api/models/EntPatient';
import { EntNurse } from '../../api/models/EntNurse';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntRoom } from '../../api/models/EntRoom';
import SaveIcon from '@material-ui/icons/Save';
import Swal from 'sweetalert2';
import Avatar from '@material-ui/core/Avatar';
import { deepOrange, deepPurple } from '@material-ui/core/colors';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            margin: theme.spacing(2),
        },
        root2: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            margin: theme.spacing(0.01),
        },
        root3: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            margin: theme.spacing(0.001),
        },
        root4: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            margin: theme.spacing(0),
        },
        root5: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            margin: theme.spacing(2),
        },
        margin: {
            margin: theme.spacing(0.5),
        },
        withoutLabel: {
            marginTop: theme.spacing(3),
        },
        textField: {
            width: '25ch',
        },
        formControl: {
            width: 500,
        },
        formDate: {
            width: 250,
        },
        formNum: {
            width: 120,
        },
        orange: {
            color: theme.palette.getContrastText(deepOrange[500]),
            backgroundColor: deepOrange[500],
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
    const [dateappoint, setDateAppoint] = useState("");
    const [timeappoint, setTimeAppoint] = useState("");
    const [addtimesave, setAddtimeSave] = useState('');

    const [cause, setCause] = React.useState(String);
    const [advice, setAdvice] = React.useState(String);
    const [hour, setHour] = React.useState(String);
    const [minute, setMinute] = React.useState(String);

    //error
    const [causeError, setCauseError] = React.useState(String);
    const [adviceError, setAdviceError] = React.useState(String);
    const [hourError, setHourError] = React.useState(String);
    const [minuteError, setMinuteError] = React.useState(String);

    const [loading, setLoading] = useState(true);
    //get data to ui materials
    useEffect(() => {  

        var setformat = moment().format();
        var setformat2 = moment('gibberish').format('YYYY MM DD');  

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
        setAddtimeSave(setformat);
        setDateAppoint(setformat2);
        setTimeAppoint(setformat2);
        setLoading(false);
    }, [loading]);

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
    const dateappointHandleChange = (event: any) => {
        CheckPattern("date",event.target.value as string);
        setDateAppoint(event.target.value as string);
    };
    const timeappointHandleChange = (event: any) => {
        CheckPattern("time",event.target.value as string);
        setTimeAppoint(event.target.value as string);
    };
    //error
    const causeHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setCause(event.target.value as string);
        CheckPattern("cause",event.target.value as string);
      };
    const adviceHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAdvice(event.target.value as string);
        CheckPattern("advice",event.target.value as string);
    };
    const hourHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setHour(event.target.value as string);
        CheckPattern("hour",event.target.value as string);
      };
    const minuteHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
        setMinute(event.target.value as string);
        CheckPattern("minute",event.target.value as string);
    };



    //checkValue
    const checkCause = (value: string) =>{
        return value.length>0;
    }
    const checkAdvice = (value: string) =>{
        return value.length>0;
    }
    const checkHour = (value: string) =>{
        return (Number(value)<4&&Number(value)>0||Number(value)==0)&&value.length>0;
    }
    const checkMinute = (value: string) =>{
        return (Number(value)<60&&Number(value)>0||Number(value)==0)&&value.length>0;
    }

    //CheckPattern
    const CheckPattern = (id: string, value:string) => {
        switch(id){
            case 'cause' :
                checkCause(value)? setCauseError(''): setCauseError('กรุณากรอกสาเหตุที่นัด!');
                return;
            case 'advice' :
                checkAdvice(value)? setAdviceError(''): setAdviceError('กรุณากรอกคำแนะนำ!');
                return;
            case 'hour' :
                checkHour(value)? setHourError(''): setHourError('กรุณากรอกตัวเลขระหว่าง 0-3!');
                return;
            case 'minute' :
                checkMinute(value)? setMinuteError(''): setMinuteError('กรุณากรอกตัวเลขระหว่าง 0-59!');
                return;
            default:
                return;
        }

    }
    const Toast = Swal.mixin({
        toast: true,
        position: 'center-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        width: 350,    
        didOpen: toast => {
            toast.addEventListener('mouseenter', Swal.stopTimer);
            toast.addEventListener('mouseleave', Swal.resumeTimer);
        },
    });
//Clear Data in field
function Clear() {
    setNurseID(0);
    setPatientID(0);
    setDoctorID(0);
    setRoomID(0);
    setCause('');
    setAdvice('');
    setHour('');
    setMinute('');
    setDateAppoint('');
    setTimeAppoint('');
}


   function Create() {

    CheckPattern("cause",cause);
    CheckPattern("advice",advice);
    CheckPattern("hour",hour);
    CheckPattern("minute",minute);

    const apiUrl = 'http://localhost:8080/api/v1/appointmentresultss';
    const appointmentresults = {
        Patient :        patientID,
        Nurse :          nurseID,  
        Doctor :         doctorID, 
        Room :           roomID,
        Cause :          cause,
        Advice :         advice,
        Hour :           hour || 0,
        Minute :         minute || 0,
        DateAppoint :    (dateappoint + "T00:00:00+07:00")   ,
        TimeAppoint :    "2000-01-01T" + timeappoint + ":00+07:00",
        AddtimeSave :    addtimesave
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
    const aleartMessage = (icon:any, title:any) =>{
        Toast.fire({
            icon: icon,
            title: title,
        });
    }

    const checkCaseError = (field: string) =>{
        switch(field) {
            case 'causeAppoint':
                aleartMessage("error","กรุณากรอกสาเหตุการนัด!");
                return;
            case 'advice':
                aleartMessage("error","กรุณากรอกคำแนะนำ!");
                return;
            case 'hourBeforeAppoint':
                aleartMessage("error","จำนวนชั่วโมงยื่นบัตรนัดไม่ถูกต้อง!");
                return;
            case 'minuteBeforeAppoint':
                aleartMessage("error","จำนวนนาทียื่นบัตรนัดไม่ถูกต้อง!");
                return;
            default:
                aleartMessage("error","บันทึกข้อมูลไม่สำเร็จ!");
                return;
        }
    }

    return ( 
        <Page theme={pageTheme.home}>
            
            <Header
                title={`ระบบนัดหมาย`}
                subtitle={`Appointment System`}
            >
             
             <Avatar className={classes.orange} style={{ height: 65, width: 65 }} >N</Avatar>
              
                <Link component={RouterLink} to='/loginappointment'>
                    <Button 
                        style={{ marginLeft: 15 }}
                        variant="contained"
                        size="large"
                     >
                            <b>ออกจากระบบ</b>
                    </Button>
                </Link>       

            </Header>

            <Content>
            <ContentHeader title="สร้างรายการนัดหมาย" >
                <Link component={RouterLink} to='/searchappointment'>
                <Button
                            href="/searchappointment"
                            variant="contained"
                            color="primary"
                            startIcon={<SearchTwoToneIcon />}
                        >
                            ระบบค้นหาข้อมูลการนัดหมาย
                </Button>
                </Link> 
            </ContentHeader>

            
                <div className={classes.root}>
                                           
                    <FormControl
                            fullWidth
                            className={classes.formControl}
                            variant="outlined"
                            
                        >
                            <InputLabel>ผู้ใช้ระบบ</InputLabel>
                            <Select
                                id="nurses"
                                label="Nurse"
                                name="nurse"
                                value={nurseID || ''}
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
                    </div>
                    <div className={classes.root}>
                        <FormControl
                            fullWidth
                            className={classes.formControl}
                            variant="outlined"
                        >
                            <InputLabel>เลือกผู้ป่วย</InputLabel>
                            <Select
                                id="patients"
                                label="Patient"
                                name="patient"
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
                    </div>
                    <div className={classes.root}>
                        <FormControl
                            fullWidth
                            className={classes.formControl}
                            variant="outlined"
                        >
                            <InputLabel>เลือกแพทย์</InputLabel>
                            <Select
                                id="doctors"
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
                    </div>
                    <div className={classes.root}>
                        <FormControl
                            fullWidth
                            className={classes.formControl}
                            variant="outlined"
                        >
                            <InputLabel>เลือกห้องตรวจ</InputLabel>
                            <Select
                                id="rooms"
                                label="Room"
                                name="room"
                                value={roomID || ''}
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
                    </div>   
                    <div className={classes.root}>
                        <FormControl
                            fullWidth
                            className={classes.formControl}
                            variant="outlined"
                        >
                            <TextField  
                                error = {causeError? true:false}
                                id="outlined-basic" 
                                label="สาเหตุที่นัด" 
                                variant="outlined" 
                                helperText={causeError}
                                value={cause}
                                onChange={causeHandle}
                            />
                            
                        </FormControl>
                    </div>   
                    <div className={classes.root}>
                        <FormControl
                            fullWidth
                            className={classes.formControl}
                            variant="outlined"
                        >
                            <TextField  
                                error = {adviceError? true:false}
                                id="outlined-basic" 
                                label="คำแนะนำก่อนพบแพทย์" 
                                variant="outlined" 
                                helperText={adviceError}
                                value={advice}
                                onChange={adviceHandle}
                                />
                        </FormControl>
                    </div>   


                        <div className={classes.root2}>
                        
                            <TableCell>
                            กำหนดยื่นบัตรนัดล่วงหน้า :
                            <TextField className={classes.formNum}
                                error = {hourError? true:false}
                                id="outlined-number"
                                label=""
                                type="number"
                                size="small"
                                helperText={hourError}
                                value={hour}
                                onChange={hourHandle}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                variant="outlined"
                                />ชั่วโมง
                            
                            </TableCell>

                            <TableCell>
                            
                            <TextField className={classes.formNum}
                                error = {minuteError? true:false}
                                id="outlined-number"
                                label=""
                                type="number"
                                size="small"
                                helperText={minuteError}
                                value={minute}
                                onChange={minuteHandle}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                variant="outlined"
                                />นาที
                            
                            </TableCell>                                                    
                        </div>


                        <div className={classes.root2}>
                                                                                                                                   
                            <TextField className={classes.formDate}
                                disabled
                                id="datetime-local"
                                label="วัน/เวลาที่บันทึก"
                                value={addtimesave}
                                onChange={addtimesaveHandleChange}
                                InputLabelProps={{
                                shrink: true
                                }}
                            />
                            
                            <TextField className={classes.formDate}
                                id="datetime-local"
                                label="วันนัดหมาย"
                                type="date" 
                                value={dateappoint || ''}
                                onChange={dateappointHandleChange}
                                InputLabelProps={{
                                    shrink: true
                                }}
                            />      
                                                                                         
                        </div>
                        
                        <div className={classes.root2}>
                        <TextField className={classes.formDate}
                                id="time"
                                label="เวลานัดหมาย"
                                type="time" 
                                value={timeappoint || ''}
                                onChange={timeappointHandleChange}
                                InputLabelProps={{
                                    shrink: true
                                }}
                            />                                                                                                   
                                                                                                                                                                                         
                        </div>


                        <div className={classes.root5}>
                        <TableRow>
                            <TableCell>
                            
                                <Button className={classes.formNum}
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
                                <Button className={classes.formNum}
                                    variant="outlined" 
                                    size="large"
                                    color="secondary"
                                    >
                                        ผลลัพธ์
                                </Button>
                            </Link>
                           
                            </TableCell>                                                      
                            </TableRow>
                        </div>
                        
                    
                        
                    
                
                <TableCell></TableCell>
            </Content>
        </Page>
    );
}

export default CreateAppointmentResults;