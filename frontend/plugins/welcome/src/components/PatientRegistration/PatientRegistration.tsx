import React, { useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  TableCell,
  Typography,
} from '@material-ui/core';

import { DefaultApi } from '../../api/apis';
import moment from 'moment';
import Swal from 'sweetalert2';

import { EntPrefix } from '../../api/models/EntPrefix';
import { EntGender } from '../../api/models/EntGender';
import { EntBloodType } from '../../api/models/EntBloodType';

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

export default function Create() {

  const classes = useStyles();
  const api = new DefaultApi();

  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [prefixs, setPrefixs] = React.useState<EntPrefix[]>([]);
  const [bloodtypes, setBloodTypes] = React.useState<EntBloodType[]>([]);

  const [userName, setName] = useState(String);
  const [genderID, setGenderID] = React.useState(Number);
  const [prefixID, setPrefixID] = React.useState(Number);
  const [bloodtypeID, setBloodTypeID] = React.useState(Number);
  const [personalID, setPersonalID] = React.useState(String);
  const [hospitalNumber, setHospitalNumber] = React.useState(String);
  const [patientName, setPatientName] = React.useState(String);
  const [drugAllergy, setDrugAllergy] = React.useState(String);
  const [mobileNumber, setMobileNumber] = React.useState(String);
  const [added, setadded] = useState('');

  const [hospitalNumberError, setHospitalNumberError] = React.useState('');
  const [personalIDError, setPersonalIDError] = React.useState('');
  const [mobileNumberError, setMobileNumberError] = React.useState('');

  const [loading, setLoading] = useState(true);


  //getdataToCombobox and checkLogin
  useEffect(() => {
    const checkUser = async () => {
      const user = JSON.parse(String(localStorage.getItem('medicalrecord-email')));
      const name = JSON.parse(String(localStorage.getItem('medicalrecord-name')));
      setLoading(false);
      const check1 = 'wuttisak@gmail.com';
      const check2 = 'aun@gmail.com';
      if (
        user != check1 &&
        user != check2
      ) {
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
          history.pushState('', '', './medicalrecordsignin');
          window.location.reload(false);
        }, 3000);
      }
      setName(name)
    };
    checkUser();

    const getGenders = async () => {
      const res = await api.listGender({ limit: 1000, offset: 0 });
      setGenders(res);
    };

    const getPrefixs = async () => {
      const res = await api.listPrefix({ limit: 1000, offset: 0 });
      setPrefixs(res);
    };

    const getBloodTypes = async () => {
      const res = await api.listBloodtype({ limit: 1000, offset: 0 });
      setBloodTypes(res);
    };

    var date = moment().format();

    setadded(date);
    getGenders();
    getPrefixs();
    getBloodTypes();
    setLoading(false);
  }, [loading]);

  // ฟังก์ชั่นสำหรับ validate เลขประจำตัวประชาชน
  const validatePersonalID = (val: string) => {
    return val.length == 13 ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate หมายเลขผู้ป่วย
  const validateHospitalNumber = (val: string) => {
    return val.match("HN\\d");
  }

  // ฟังก์ชั่นสำหรับ validate เบอร์โทรศัพท์
  const validateMobileNumber = (val: string) => {
    return val.length == 10 ? true : false;
  }

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'hospitalNumber':
        validateHospitalNumber(value) ? setHospitalNumberError('') : setHospitalNumberError('รหัสนักศึกษาขึ้นต้นด้วย HN ตามด้วยตัวเลข');
        return;
      case 'personalID':
        validatePersonalID(value) ? setPersonalIDError('') : setPersonalIDError('เลขประจำตัวประชาชน 13 หลัก');
        return;
      case 'mobileNumber':
        validateMobileNumber(value) ? setMobileNumberError('') : setMobileNumberError('เบอร์โทรศัพท์ 10 หลัก')
        return;
      default:
        return;
    }
  }

  const checkCaseSaveError = (field: string) => {
    switch (field) {
      case 'hospitalNumber':
        alertMessage("error", "รหัสนักศึกษาขึ้นต้นด้วย HN ตามด้วยตัวเลข");
        return;
      case 'personalID':
        alertMessage("error", "เลขประจำตัวประชาชน 13 หลัก");
        return;
      case 'mobileNumber':
        alertMessage("error", "เบอร์โทรศัพท์ 10 หลัก");
        return;
      default:
        alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  //handle
  const genderHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGenderID(event.target.value as number);
  };
  const prefixHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPrefixID(event.target.value as number);
  };
  const bloodtypeHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBloodTypeID(event.target.value as number);
  };
  const personalIDHandle = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('personalID', validateValue)
    setPersonalID(event.target.value as string);
  };
  const hospitalNumberHandle = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('hospitalNumber', validateValue)
    setHospitalNumber(event.target.value as string);
  };
  const patientNameHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientName(event.target.value as string);
  };
  const drugAllergyHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDrugAllergy(event.target.value as string);
  };
  const mobileNumberHandle = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('mobileNumber', validateValue)
    setMobileNumber(event.target.value as string);
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
    setGenderID(0);
    setPrefixID(0);
    setBloodTypeID(0);
    setPersonalID('');
    setHospitalNumber('');
    setPatientName('');
    setDrugAllergy('');
    setMobileNumber('');
  }

  //create
  const Create = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const patient = {
      PersonalID: personalID,
      HospitalNumber: hospitalNumber,
      Prefix: prefixID,
      PatientName: patientName,
      Gender: genderID,
      BloodType: bloodtypeID,
      DrugAllergy: drugAllergy,
      MobileNumber: mobileNumber,
      Added: added,
    };
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patient),
    };

    console.log(patient);

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
          checkCaseSaveError(data.error.Name);
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>

      <Header
        title="ระบบลงทะเบียนผู้ป่วย"
        subtitle="สามารถลงทะเบียนผู้ป่วยได้ที่นี่"
      >
        <Typography>
          {userName}
        </Typography>
        &nbsp;&nbsp;&nbsp;
        <Button
          component={RouterLink}
          to="/medicalrecordsignin"
          variant="contained"
          color="secondary"
        >
          Logout
          </Button>
      </Header>

      <Content>
        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                error={personalIDError ? true : false}
                id="personalID"
                label="เลขประจำตัวประชาชน"
                name="PersonalID"
                variant="outlined"
                type="string"
                size="medium"
                inputProps={{ maxLength: 13 }}
                helperText={personalIDError}
                value={personalID}
                onChange={personalIDHandle}
              />
            </FormControl>

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                error={hospitalNumberError ? true : false}
                id="hospitalNumber"
                label="หมายเลขผู้ป่วย"
                name="HospitalNumber"
                variant="outlined"
                type="string"
                size="medium"
                helperText={hospitalNumberError}
                value={hospitalNumber}
                onChange={hospitalNumberHandle}
              />
            </FormControl>

            <FormControl
              className={classes.margin}
              fullWidth
              variant="outlined"
            >
              <InputLabel>คำนำหน้า</InputLabel>
              <Select
                id="prefixs"
                label="คำนำหน้า"
                name="prefix"
                value={prefixID}
                onChange={prefixHandle}
              >
                {prefixs.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.prefixValue}
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
                id="patientName"
                label="ชื่อ - สกุล"
                name="PatientName"
                variant="outlined"
                type="string"
                size="medium"
                value={patientName}
                onChange={patientNameHandle}
              />
            </FormControl>

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel>เพศ</InputLabel>
              <Select
                id="genders"
                label="เพศ"
                name="gender"
                value={genderID}
                onChange={genderHandle}
              >
                {genders.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.genderValue}
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
              <InputLabel>หมู่เลือด</InputLabel>
              <Select
                id="bloodtypes"
                label="หมู่เลือด"
                name="bloodtype"
                value={bloodtypeID}
                onChange={bloodtypeHandle}
              >
                {bloodtypes.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.bloodValue}
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
                id="drugAllergy"
                label="ประวัติการแพ้ยา"
                name="DrugAllergy"
                variant="outlined"
                type="string"
                size="medium"
                value={drugAllergy}
                onChange={drugAllergyHandle}
              />
            </FormControl>

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                error={mobileNumberError ? true : false}
                id="mobileNumber"
                label="เบอร์โทรศัพท์"
                name="MobileNumber"
                variant="outlined"
                type="string"
                size="medium"
                inputProps={{ maxLength: 10 }}
                helperText={mobileNumberError}
                value={mobileNumber}
                onChange={mobileNumberHandle}
              />
            </FormControl>

            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                id="datetime-local"
                label="วันที่ลงทะเบียน"
                name="added"
                variant="outlined"
                size="medium"
                value={added}
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
                  ลงทะเบียน
                                </Button>
              </TableCell>

              <TableCell>
                <Button
                  component={RouterLink}
                  to='/patientregistrationtable'
                  variant="contained"
                  size="large"
                  color="secondary"
                >
                  ผลการลงทะเบียน
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
