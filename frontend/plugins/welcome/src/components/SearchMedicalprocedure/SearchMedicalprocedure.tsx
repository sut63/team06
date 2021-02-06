import React, { useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button, Grid } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import Swal from 'sweetalert2'
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { Cookies } from '../../Cookie';
import { EntPatient } from '../../api/models/EntPatient';
import { EntMedicalProcedure } from '../../api/models/EntMedicalProcedure';


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
        },
        margin: {
            margin: theme.spacing(2),
        },
        insideLabel: {
            margin: theme.spacing(1),
        },
        button: {
            marginLeft: '50px',
        },
        textField: {
            width: 600,
            marginLeft: 10,
            marginRight: -10,
        },
        select: {
            width: 600,
            marginLeft: 10,
            marginRight: -10,
        },
        paper: {
            marginTop: theme.spacing(1),
            marginBottom: theme.spacing(1),
            marginLeft: theme.spacing(1),
        },
        pa: {
            marginTop: theme.spacing(2),
        },
        table: {
            minWidth: 650,
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


export default function ComponentsTable() {
    const classes = useStyles();
    const http = new DefaultApi();
    var cookie = new Cookies();
    var nursename = cookie.GetCookie("nursename");
    const [loading, setLoading] = useState(true);
    const [search, setSearch] = useState(false);

    const [checkpatient, setPatientnames] = useState(false);
    const [patients, setPatients] = useState<EntPatient[]>([])
    const [medical, setmedical] = useState<EntMedicalProcedure[]>([])

    const [patientname, setPatientname] = useState(String);
    const alertMessage = (icon: any, title: any) => {
        Toast.fire({
            icon: icon,
            title: title,
        });
        setSearch(false);
    }

    useEffect(() => {
        const getPatients = async () => {
            const res = await http.listPatient({ offset: 0 });
            setLoading(false);
            setPatients(res);
        };
        getPatients();
        const getMedicalprocedures = async () => {
            const res = await http.listMedicalprocedure({ offset: 0 });
            setLoading(false);
            setmedical(res);
        };
        getMedicalprocedures();
        if (nursename == "") {
            Swal.fire({
                title: 'โปรดเข้าสู่ระบบก่อนเข้าใช้งาน',
                position: 'center',
                showConfirmButton: true,
                timer: 6000,
                timerProgressBar: false,
                didOpen: toast => {
                    toast.addEventListener('mouseenter', Swal.stopTimer);
                    toast.addEventListener('mouseleave', Swal.resumeTimer);
                },
            });
            setTimeout(() => {
                window.location.replace("http://localhost:3000/loginmedicalprocedure");
            }, 3000);
        }
    }, [loading]);

    const patientnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setSearch(false);
        setPatientnames(false);
        setPatientname(event.target.value as string);

    };

    const cleardata = () => {
        setPatientname("");
        setSearch(false);
        setPatientnames(false);
        setSearch(false);

    }

    const checkresearch = async () => {
        var check = false;
        medical.map(item => {
            if (patientname != "") {
                if (item.edges?.patient?.patientName?.startsWith(patientname)) {
                    setPatientnames(true);
                    alertMessage("success", "ค้นหาข้อมูลสำเร็จ");
                    check = true;
                }
            }
        })
        if (!check) {
            alertMessage("error", "ไม่พบข้อมูล");
        }
        //console.log(checkpersonalID)
        if (patientname == "") {
            alertMessage("info", "แสดงข้อมูลในระบบ");
        }
    };

    return (

        <Page theme={pageTheme.home}>
            <Header title="ระบบค้นหาการทำหัตถการ">
                <TableCell align={
                    "center"} >
                    <text><strong> {nursename}</strong></text>
                    <br /><br />
                    <Button
                        style={{ backgroundColor: '#4fc3f7	' }}
                        href="/medicalprocedure"
                        variant="contained"
                        color="secondary"
                    >
                        BACK
                  </Button>
                </TableCell>
            </Header>
            <Content>

                <div className={classes.root}>
                    <div ><form noValidate autoComplete="off">
                        <FormControl
                            fullWidth
                            className={classes.margin}
                            variant="outlined"
                            size="medium"
                        >
                            <TextField
                                style={{ width: 500, marginLeft: 8, marginRight: -7, marginTop: 5 }}
                                className={classes.textField}
                                id="personalID"
                                variant="outlined"
                                label="กรอกชื่อผู้ป่วย"
                                color="primary"
                                type="string"
                                size="medium"
                                value={patientname}
                                onChange={patientnamehandlehange}
                            />
                        </FormControl>
                    </form>

                    </div>
                </div>
                <div className={classes.root}>
                    <Button
                        style={{ backgroundColor: '#e57373	' }}
                        onClick={() => {
                            checkresearch();
                            setSearch(true);
                        }}
                        variant="contained"
                        size="large"
                        startIcon={<SearchTwoToneIcon />}
                    >
                        search
               </Button>
                </div>

                <br></br>

                <Grid container justify="center">
                    <Grid item xs={12} md={10}>
                        <Paper>
                            {search ? (
                                <div>
                                    {  checkpatient ? (
                                        <TableContainer component={Paper}>
                                            <Table className={classes.table} aria-label="simple table">
                                                <TableHead>
                                                    <TableRow>
                                                        <TableCell align="center">ชื่อออเดอร์</TableCell>
                                                        <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                                        <TableCell align="center">ประเภทการทำหัตถการ</TableCell>
                                                        <TableCell align="center">ชื่อหมอ</TableCell>
                                                        <TableCell align="center">คำอธิบาย</TableCell>
                                                        <TableCell align="center">วันที่-เวลา</TableCell>

                                                    </TableRow>
                                                </TableHead>
                                                <TableBody>

                                                    {medical.filter((filter: any) => filter.edges?.patient?.patientName.startsWith(patientname)).map((item: any) => (
                                                        <TableRow key={item.id}>
                                                            <TableCell align="center">{item.procedureOrder}</TableCell>
                                                            <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                                                            <TableCell align="center">{item.edges?.procedureType?.procedureName}</TableCell>
                                                            <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                                                            <TableCell align="center">{item.procedureDescripe}</TableCell>
                                                            <TableCell align="center">{moment(item.addtime).format('LLLL')}</TableCell>
                                                        </TableRow>
                                                    ))}
                                                </TableBody>
                                            </Table>
                                        </TableContainer>
                                    )
                                        : patientname == "" ? (
                                            <div>
                                                <TableContainer component={Paper}>
                                                    <Table className={classes.table} aria-label="simple table">
                                                        <TableHead>
                                                            <TableRow>
                                                                <TableCell align="center">ชื่อออเดอร์</TableCell>
                                                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                                                <TableCell align="center">ประเภทการทำหัตถการ</TableCell>
                                                                <TableCell align="center">ชื่อหมอ</TableCell>
                                                                <TableCell align="center">คำอธิบาย</TableCell>
                                                                <TableCell align="center">วันที่-เวลา</TableCell>
                                                            </TableRow>
                                                        </TableHead>
                                                        <TableBody>

                                                            {medical.map((item: any) => (
                                                                <TableRow key={item.id}>
                                                                    <TableCell align="center">{item.procedureOrder}</TableCell>
                                                                    <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                                                                    <TableCell align="center">{item.edges?.procedureType?.procedureName}</TableCell>
                                                                    <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
                                                                    <TableCell align="center">{item.procedureDescripe}</TableCell>
                                                                    <TableCell align="center">{moment(item.addtime).format('LLLL')}</TableCell>
                                                                </TableRow>
                                                            ))}
                                                        </TableBody>
                                                    </Table>
                                                </TableContainer>

                                            </div>
                                        ) : null}
                                </div>
                            ) : null}
                        </Paper>
                    </Grid>
                </Grid>
            </Content>
        </Page>
    );
}