import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
    MuiPickersUtilsProvider,
    KeyboardTimePicker,
    KeyboardDatePicker,
} from '@material-ui/pickers';
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
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDiagnosis, EntPatient } from '../../api';
import { Cookies } from '../../Cookie'
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
// header css
const HeaderCustom = {
    minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        marginTop: theme.spacing(2),
        marginBottom: theme.spacing(2),
    },
    formControl: {
        width: 300,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        width: 300,
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
}));


const Searchdiagnosis: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()

    // Customer
    const [idPatient, setIDPatient] = React.useState<number>(0)
    const [patient, setPatient] = React.useState<EntPatient[]>([])
    const getpatient = async () => {
        const res = await api.listPatient({})
        setPatient(res)
    }

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


    // CheckIn  
    var lenDiagnosis: number
    const [diagnosis, setDiagnosis] = React.useState<EntDiagnosis[]>([])
    const getDiagnosiss = async () => {
        const res = await api.getDiagnosis({ id: idPatient })
        setDiagnosis(res)
        lenDiagnosis = res.length
        if (lenDiagnosis > 0) {
            setOpen(true)
        } else {
            setFail(true)
        }
    }




    // set data to object and setIdcustomer
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const name = event.target.name as keyof typeof Searchdiagnosis;
        const { value } = event.target;
        setIDPatient(value);
    };

    // clear cookies
    function Clears() {
        ck.ClearCookie()
        window.location.reload(false)
    }

    // function seach data
    function seach() {
        getDiagnosiss();
    }
    // Lifecycle Hooks
    useEffect(() => {
        getpatient();
    }, []);




    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาการตรวจวินิจฉัย`}>
                <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
                <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
                <Button
                    style={{ marginLeft: 20 }}
                    component={RouterLink}
                    to="/diagnosis"
                    variant="contained"
                >
                    ย้อนกลับ
                    &nbsp;
             </Button>
            </Header>
            <Content>
                <Grid container spacing={1}>
                    <Grid item xs={1}>
                        <div className={classes.paper}><h3>ชื่อผู้ป่วย</h3></div>
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel> เลือกผู้ป่วย</InputLabel>
                            <Select
                                id="Patient"
                                name="Patient"
                                value={idPatient || ''} // (undefined || '') = ''
                                onChange={handleChange}
                            >
                                {patient.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.patientName}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <Button
                            variant="contained"
                            color="primary"
                            size="large"
                            onClick={seach}
                        >
                            ค้นหา
                        </Button>
                    </Grid>
                </Grid>

                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ตารางการตรวจวินิจฉัย:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">ชื่อแพทย์</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย </TableCell>
                                <TableCell align="center">ประเภทการรักษา</TableCell>
                                <TableCell align="center">อาการ</TableCell>
                                <TableCell align="center">ความคิดเห็นจากแพทย์</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>

                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {diagnosis.map((item: EntDiagnosis) => (
                                <TableRow key={item.id}>

                                    <TableCell align="center">{item.edges?.doctorName?.doctorName}</TableCell>
                                    <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                                    <TableCell align="center">{item.edges?.type?.type}</TableCell>
                                    <TableCell align="center">{item.symptom}</TableCell>
                                    <TableCell align="center">{item.opinionresult}</TableCell>
                                    <TableCell align="center">{item.note}</TableCell>

                                    <TableCell align="center">

                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="success">
                        ค้นหาสำเร็จ
          </Alert>
                </Snackbar>

                <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="error">
                        ไม่พบข้อมูล
          </Alert>
                </Snackbar>
            </Content>
        </Page>
    );
};

export default Searchdiagnosis;