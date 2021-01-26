import React, { useEffect, useState } from 'react';

import Box from '@material-ui/core/Box';
import { Link as RouterLink } from 'react-router-dom';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatient } from '../../api/models/EntPatient';
import { EntTreatmentType } from '../../api/models/EntTreatmentType';
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
import { Cookies } from '../../Cookie'
import Avatar from '@material-ui/core/Avatar';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
    Link,
    identityApiRef,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';
import { MenuItem } from '@material-ui/core';
import Icon from '*.icon.svg';
import Swal from 'sweetalert2';

interface Diagnosis {
    TreatmentType: number;
    Patient: number;
    Doctor: number;
    Symptom: string;
    Opinionresult: string;
    diagnosisDate: string;
    note: string;
    // create_by: number;
}

const HeaderCustom = {
    minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
        },
        margin: {
            margin: theme.spacing(3),
        },
        withoutLabel: {
            marginTop: theme.spacing(3),
        },
        textField: {
            width: '25ch',
        },
        formControl: {
            width: 300,
        },
        
        paper: {
            marginTop: theme.spacing(2),
            marginBottom: theme.spacing(2),
        },
    
        selectEmpty: {
            marginTop: theme.spacing(2),
        },
        container: {
            display: 'flex',
            flexWrap: 'wrap',
        },
       
        bottom: {
            marginLeft: theme.spacing(3),
            marginTop: theme.spacing(1),
            marginBottom: theme.spacing(2),
        },
        divider: {
            margin: theme.spacing(2, 0),
        },
        table: {
            minWidth: 650,
        },
    }),
);

export default function Create() {
    const classes = useStyles();
    const profile = { givenName: 'ระบบวินิจฉัย' };
    const api = new DefaultApi();

    const [Diagnosis, setDiagnosisitem] = React.useState<Partial<Diagnosis>>({});
    const [TreatmentTypeID, SetType] = React.useState<EntTreatmentType[]>([]);
    const [PatientID, setPatient] = React.useState<EntPatient[]>([]);
    const [DoctorID, Setdoctor] = React.useState<EntDoctor[]>([]);

    const [identifiCationSymptomError, setidentifiCationSymptomError] = React.useState('');
    const [identifiNoteError, setIdentifiNoteError] = React.useState('');
    const [identifiOpinionresultError, setOpinionresultError] = React.useState('');


    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()
    var cookieID = ck.GetID()


    // alert setting
    const [open, setOpen] = React.useState(false);
    const [fail, setFail] = React.useState(false);
    const [errors, setError] = React.useState(String);

    const checkCaseSaveError = (s: string) => {
        switch (s) {
            case 'symptom':
                setError("กรุณากรอกอาการให้ครบ")
                return;
            case 'Opinionresult':
                setError("กรุณากรอกความเห็นให้ครบ")
                return;
            case 'note':
                setError("กรอกหมายเหตุให้ครบ")
                return;
            default:
                return;
        }
    };

    //close alert 
    const handleClose = (_event?: React.SyntheticEvent, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setFail(false);
        setOpen(false);
    };

    const getTreatmentType = async () => {
        const res = await api.listTreatmentType({ limit: 10, offset: 0 });
        SetType(res);
    };
    const getdoctor = async () => {
        const res = await api.listDoctor({ limit: 10, offset: 0 });
        Setdoctor(res);
    };

    const getpatient = async () => {
        const res = await api.listPatient({ limit: 10, offset: 0 });
        setPatient(res);
    };







    /*const Createsection = async () => {
      const res: any = await api.createSection({ section });
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
      } else {
        setAlert(false);
      }
    };*/

    /*const handleChanges = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
        //const name = event.target.name as keyof typeof Create;
        //const { value } = Number(cookieID)
        //const validateValue = value.toString()
        //checkPattern(name, validateValue)
        setDiagnosisitem({ ...Diagnosis, [Diagnosis.Doctor]: Number(cookieID) });
        console.log(Diagnosis);
    };
*/
    const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
        const name = event.target.name as keyof typeof Create;
        const { value } = event.target;
        const validateValue = value.toString()
        checkPattern(name, validateValue)
        setDiagnosisitem({ ...Diagnosis, [name]: value });
        console.log(Diagnosis);
    };

    // list Doctor
    const [doctor, setDoctor] = React.useState<EntDoctor>()
    const listDoctor = async () => {
        const res = await api.getDoctor({ id: Number(cookieID) })
        setDoctor(res)
        console.log(doctor?.doctorName);

    }


    const alertMessage = (icon: any, title: any) => {
        Toast.fire({
            icon: icon,
            title: title,
        });
    }

    // Alert setting
    const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: Toast => {
            Toast.addEventListener('mouseenter', Swal.stopTimer);
            Toast.addEventListener('mouseleave', Swal.resumeTimer);
        }
    })

    //ฟังก์ชัน validate อาการ
    const ValidateNoteError = (val: string) => {
        return val.length > 1 ? true : false;
    }

    //ฟังก์ชัน validate อาการ
    const ValidateOpinionresultError = (val: string) => {
        return val.length > 1 ? true : false;
    }

    //ฟังก์ชัน validate อาการ
    const ValidateCationSymptomError = (val: string) => {
        return val.length > 1 ? true : false;
    }

    // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
    const checkPattern = (name: string, value: string) => {
        switch (name) {
            case 'Symptom':
                ValidateCationSymptomError(value) ? setidentifiCationSymptomError('') : setidentifiCationSymptomError('กรุณาระบุอาการให้ครบ');
                return;
            case 'note':
                ValidateNoteError(value) ? setIdentifiNoteError('') : setIdentifiNoteError('ระบุหมายเหตุให้ครบ');
                return;
            case 'Opinionresult':
                ValidateOpinionresultError(value) ? setOpinionresultError('') : setOpinionresultError('กรุณาระบุความเห็นจากแพทย์ให้ครบ')
                return;
            default:
                return;
        }
    }

    function save() {
        const apiUrl = 'http://localhost:8080/api/v1/Diagnosiss';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(Diagnosis),
        }

        console.log(Diagnosis); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status == true) {
                    Toast.fire({
                        icon: 'success',
                        title: 'บันทึกสำเร็จ'
                    })
                    clear();
                    setOpen(true);
                } else {
                    setFail(true);
                    checkCaseSaveError(data.error.Name)
                }
            })
    }

    function Clears() {
        ck.ClearCookie()
        window.location.reload(false)
    }

    //set time
    const [diagnosis_Date, setdiagnosisDate] = React.useState<any>(0)

    // set data for time in
    const handleChangediagnosisDate = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const name = event.target.name as keyof typeof Create;
        const { value } = event.target;
        const time = value + ":00+07:00"
        setDiagnosisitem({ ...Diagnosis, [name]: time });
        setdiagnosisDate(value);
        console.log(diagnosis_Date);
    };

    // clear input form
    function clear() {
        setDiagnosisitem({});
    }

    useEffect(() => {
        setDiagnosisitem({ ...Diagnosis, ['Doctor']: doctor?.id })
    }, [doctor]);

    useEffect(() => {
        getdoctor();
        getpatient();
        getTreatmentType();
        listDoctor();
    }, []);

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาการตรวจวินิจฉัย`}>
                
                <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
                <Button
                    style={{ marginLeft: 20 }}
                    component={RouterLink}
                    to="/"
                    variant="contained"
                >
                    ออกจากระบบ
                    &nbsp;
             </Button>
            </Header>
            <Content>

                <ContentHeader title="ระบบการตรวจรักษา">

                    &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
          
                    <Link component={RouterLink} to="" >
                       
             &nbsp;
             <Button
                                style={{ marginRight: 600  }}

                                size="large"
                                color="inherit"
                                component={RouterLink}
                                to="/SearchdiagnosisPage"
                                variant="outlined"
                            >
                                ค้นหาการตรวจวินิจฉัย
                                
             </Button>
             &nbsp; 
        </Link>
                </ContentHeader>
                <div className={classes.root}>
                    <form noValidate autoComplete="off">
                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                        >

                            <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                                <Alert onClose={handleClose} severity="success">
                                    บันทึกสำเร็จ !
        </Alert>
                            </Snackbar>

                            <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
                                <Alert onClose={handleClose} severity="error">
                                    {errors}
                                </Alert>
                            </Snackbar>

              ชื่อแพทย์
              <FormControl variant="outlined" className={classes.formControl}>
                                <TextField
                                    disabled
                                    name="Doctor"
                                    variant="outlined"
                                    value={doctor?.doctorName}
                                    onChange={handleChange}
                                />
                            </FormControl>


              ชื่อผู้ป่วย
              <TextField
                                id="patient"
                                select
                                name="Patient"
                                label="เลือกผู้ป่วย"
                                value={Diagnosis.Patient}
                                variant="outlined"
                                onChange={handleChange}
                            >
                                {PatientID.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.patientName}
                                        </MenuItem>
                                    );
                                })}
                            </TextField>
             &nbsp;
             ประเภทการรักษา
             <TextField

                                id="TreatmentType"
                                select
                                name="TreatmentType"
                                label="ประเภทการรักษา"

                                value={Diagnosis.TreatmentType}
                                variant="outlined"
                                onChange={handleChange}
                            >
                                {TreatmentTypeID.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.type}
                                        </MenuItem>
                                    );
                                })
                                }
                            </TextField>



                        </FormControl>


                        <form className={classes.root} noValidate autoComplete="off">

                            <TextField id="outlined-basic" label="ความคิดเห็นจากแพทย์"
                                error={identifiOpinionresultError ? true : false}
                                helperText={identifiOpinionresultError}
                                variant="outlined"
                                name="Opinionresult"
                                value={Diagnosis.Opinionresult}
                                onChange={handleChange} />

                        </form>
    &nbsp;


                        <form className={classes.root} noValidate autoComplete="off">
                            <TextField id="symptom" label="อาการ"
                                variant="outlined"
                                error={identifiCationSymptomError ? true : false}
                                helperText={identifiCationSymptomError}
                                name="Symptom"
                                value={Diagnosis.Symptom}
                                onChange={handleChange} />

                        </form>
            &nbsp;
            &nbsp;
            &nbsp;

            <form className={classes.root} noValidate autoComplete="off">

                            <TextField id="outlined-basic" label="หมายเหตุ"
                                variant="outlined"
                                error={identifiNoteError ? true : false}
                                helperText={identifiNoteError}
                                name="note"
                                value={Diagnosis.note}
                                onChange={handleChange} />
                        </form>

                        <FormControl variant="outlined" className={classes.formControl}>
                            <TextField
                                name="diagnosisDate"
                                type="datetime-local"
                                value={diagnosis_Date}
                                defaultValue="2020-12-31"
                                onChange={handleChangediagnosisDate}
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

                            <Box color="text.primary" clone>
                                <Button />
                            </Box>
                        </FormControl>



                        <div className={classes.margin}>
                            <Button
                                onClick={() => {
                                    save();

                                }}
                                variant="contained"
                                color="secondary"

                            >
                                ยืนยันการลงทะเบียน
             </Button>

                       
             &nbsp;
             <Button
                                style={{ marginLeft: 20 }}
                                component={RouterLink}
                                to="/diagnosistable"
                                variant="contained"
                            >
                                ดูผลการตรวจวินิจฉัย
                                &nbsp;
             </Button>

                        </div>
                    </form>


                </div>
        หมายเหตุ
                &nbsp;

             เป็นระบบจำลองการวินิจฉัย



      </Content>

        </Page>

    );
}