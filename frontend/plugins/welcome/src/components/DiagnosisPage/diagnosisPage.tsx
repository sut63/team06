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
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';
import { MenuItem } from '@material-ui/core';




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

    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()

    // alert setting
    const [open, setOpen] = React.useState(false);
    const [fail, setFail] = React.useState(false);

    //close alert 
    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
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

    useEffect(() => {
        getdoctor();
        getpatient();
        getTreatmentType();


    }, []);

    interface Diagnosis {

        TreatmentType: number;
        Diagnosis: number;
        Patient: number;
        Doctor: number;
        Symptom: string;
        Opinionresult: string;
        diagnosisDate: string;

        // create_by: number;
    }

    /*const Createsection = async () => {
      const res: any = await api.createSection({ section });
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
      } else {
        setAlert(false);
      }
    };*/

    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
        const name = event.target.name as keyof typeof Create;
        const { value } = event.target;
        setDiagnosisitem({ ...Diagnosis, [name]: value });
        console.log(Diagnosis);
    };

    function save() {
        const apiUrl = 'http://localhost:8080/api/v1/Diagnosiss';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(Diagnosis),


        };

        console.log(Diagnosis); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data.status);
                if (data.id != null) {
                    clear();
                    setOpen(true);
                } else {
                    setFail(true);
                }
            });
    }

    function Clears() {
        ck.ClearCookie()
        window.location.reload(false)
    }

    //set time
    const [diagnosisDate, setdiagnosisDate] = React.useState<any>(0)

    // set data for time in
    const handleChangediagnosisDate = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const name = event.target.name as keyof typeof diagnosisDate;
        const { value } = event.target;
        const time = value + ":00+07:00"
        setdiagnosisDate({ ...diagnosisDate, [name]: time });
        setdiagnosisDate(value);
        console.log(diagnosisDate);
    };

    // clear input form
    function clear() {
        setDiagnosisitem({});
    }
    return (
        <Page theme={pageTheme.home}>
            <Header
                title={`Welcome ${profile.givenName || 'to Backstage'}`}
                subtitle="Some quick intro and links."

            ></Header>
            <Content>

                <ContentHeader title="ระบบการตรวจรักษา">

                    &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
          <Avatar src="/broken-image.jpg" style={{ height: 75, width: 75 }} />
                    <Link component={RouterLink} to="" >
                        <Button
                            style={{ marginLeft: 20 }}
                            component={RouterLink}
                            to="/"
                            variant="contained"
                        >
                            ออกจากระบบ
                            &nbsp;
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
                                    บันทึกไม่สำเร็จ!
        </Alert>
                            </Snackbar>

              ชื่อแพทย์
              <TextField
                                id="doctor"
                                select
                                name="DoctorID"
                                label="เลือกแพทย์"
                                value={Diagnosis.Doctor}
                                variant="outlined"
                                onChange={handleChange}
                            >
                                {DoctorID.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.doctorName}
                                        </MenuItem>
                                    );
                                })}
                            </TextField>
              ชื่อผู้ป่วย
              <TextField
                                id="patient"
                                select
                                name="PatientID"
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
                                id="course"
                                select
                                name="TreatmentTypeID"
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
                                variant="outlined"
                                name="Opinionresult"
                                value={Diagnosis.Opinionresult}
                                onChange={handleChange} />

                        </form>
    &nbsp;

    <form className={classes.root} noValidate autoComplete="off">

                            <TextField id="outlined-basic" label="อาการ"
                                variant="outlined"
                                name="Symptom"
                                value={Diagnosis.Symptom}
                                onChange={handleChange} />

                        </form>
            &nbsp;
            &nbsp;
            &nbsp;

            <FormControl variant="outlined" className={classes.formControl}>
                            <TextField
                                name="diagnosis_date"
                                type="datetime-local"
                                value={Diagnosis.diagnosisDate}
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

                            <Button
                                style={{ marginLeft: 20 }}
                                component={RouterLink}
                                to="/WelcomePage"
                                variant="contained"
                            >
                                ย้อนกลับ
                                &nbsp;
             </Button>
             &nbsp;

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

