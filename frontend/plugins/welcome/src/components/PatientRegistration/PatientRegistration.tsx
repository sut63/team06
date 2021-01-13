import React, { useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import SaveIcon from '@material-ui/icons/Save';
import AccountBoxIcon from '@material-ui/icons/AccountBox';
import Swal from 'sweetalert2';
import moment from 'moment';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';

import {
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Button,
  Box,
  TextField,
} from '@material-ui/core';

import { DefaultApi } from '../../api/apis';
import { EntPrefix } from '../../api/models/EntPrefix';
import { EntGender } from '../../api/models/EntGender';
import { EntBloodType } from '../../api/models/EntBloodType';


const useStyles = makeStyles({
  boxStyle: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    flexDirection: 'row'
  },
  formControl: {
    width: 300,
    height: 65,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
    width: 300,
    height: 65,
  },
  textField: {
    width: 300,
  },
});


export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [prefixs, setPrefixs] = React.useState<EntPrefix[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [bloodtypes, setBloodTypes] = React.useState<EntBloodType[]>([]);

  const [prefixID, setPrefixID] = React.useState(Number);
  const [genderID, setGenderID] = React.useState(Number);
  const [bloodTypeID, setBloodTypeID] = React.useState(Number);
  const [personalID, setPersonalID] = React.useState(Number);
  const [patientName, setPatientName] = React.useState(String);
  const [Age, setAge] = React.useState(Number);
  const [hospitalNumber, setHospitalNumber] = React.useState(String);
  const [drugAllergy, setDrugAllergy] = React.useState(String);
  const [addedDate, setAddedDate] = useState('');

  const [loading, setLoading] = useState(true);


  //getdataToCombobox
  useEffect(() => {
    const getPrefixs = async () => {
      const res = await api.listPrefix({ limit: 1000, offset: 0 });
      setPrefixs(res);
    };

    const getGenders = async () => {
      const res = await api.listGender({ limit: 1000, offset: 0 });
      setGenders(res);
    };

    const getBloodTypes = async () => {
      const res = await api.listBloodtype({ limit: 1000, offset: 0 });
      setBloodTypes(res);
    };

    var date = moment().format();

    setAddedDate(date);
    getPrefixs();
    getGenders();
    getBloodTypes();
    setLoading(false);
  }, [loading]);

  //handle
  const prefixHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPrefixID(event.target.value as number);
  };
  const genderHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGenderID(event.target.value as number);
  };
  const bloodtypeHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBloodTypeID(event.target.value as number);
  };
  const patientNameHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientName(event.target.value as string);
  };
  const personalIDHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonalID(event.target.value as number);
  };
  const AgeHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAge(event.target.value as number);
  };
  const hospitalNumberHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setHospitalNumber(event.target.value as string);
  };
  const drugAllergyHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDrugAllergy(event.target.value as string);
  };

  // alert setting
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
    setPrefixID(0);
    setGenderID(0);
    setBloodTypeID(0);
    setPersonalID(0);
    setPatientName('');
    setAge(0);
    setHospitalNumber('');
    setDrugAllergy('');
  }

  //Create
  const Create = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const patient = {
      personalID: personalID,
      prefix: prefixID,
      patientName: patientName,
      age: Age,
      gender: genderID,
      bloodType: bloodTypeID,
      hospitalNumber: hospitalNumber,
      drugAllery: drugAllergy,
      addedDate: addedDate,
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
        title="ระบบลงทะเบียนผู้ป่วย"
        subtitle="สามารถลงทะเบียนผู้ป่วยได้ที่นี่"
      ></Header>

      <Content>
        <ContentHeader title="ลงทะเบียนผู้ป่วย">
          <Box
            display="flex"
            justifyContent="flex-start"
            flexDirection="column"
          >
            <AccountBoxIcon style={{ fontSize: 100 }}></AccountBoxIcon>

            <Link component={RouterLink} to="/" >
              ออกจากระบบ
                        </Link>
          </Box>

        </ContentHeader>

        <Grid container>
          <Grid item xs={12}>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="personalID"
                  label="เลขประจำตัวประชาชน"
                  name="personalID"
                  variant="outlined"
                  type="number"
                  size="medium"
                  value={personalID}
                  onChange={personalIDHandle}
                />
              </FormControl>
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>คำนำหน้า</InputLabel>
                <Select
                  id="prefixs"
                  label="คำนำหน้า"
                  name="prefixValue"
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
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="patientName"
                  label="ชื่อผู้ป่วย"
                  name="patienName"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={patientName}
                  onChange={patientNameHandle}
                />
              </FormControl>
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="patientAge"
                  label="อายุ"
                  name="age"
                  variant="outlined"
                  type="number"
                  size="medium"
                  value={Age}
                  onChange={AgeHandle}
                />
              </FormControl>
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                <Select
                  id="genders"
                  label="เพศ"
                  name="genderValue"
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
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>หมู่เลือด</InputLabel>
                <Select
                  id="bloodtypes"
                  label="หมู่เลือด"
                  name="bloodValue"
                  value={bloodTypeID}
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
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="hospitalNumber"
                  label="หมายเลขผู้ป่วย"
                  name="hospitalNumber"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={hospitalNumber}
                  onChange={hospitalNumberHandle}
                />
              </FormControl>
            </Box>
            <Box className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="drugAllergy"
                  label="ประวัติการแพ้ยา"
                  name="drugAllergy"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={drugAllergy}
                  onChange={drugAllergyHandle}
                />
              </FormControl>
            </Box>
            <Box
              className={classes.boxStyle}>
              <FormControl variant="outlined" className={classes.formControl}>
              <TextField
                id="datetime-local"
                label="Date"
                name="added"
                variant="outlined"
                size="medium"
                value={addedDate}
                disabled
              />
              </FormControl>
            </Box>
            <Box display="flex" justifyContent="center" >
              <Button
                style={{ width: 140, height: 40 }}
                variant="contained"
                color="primary"
                startIcon={<SaveIcon />}
                onClick={Create}>
                บันทึกข้อมูล
                            </Button>
              <Button
                style={{ width: 140, height: 40, marginLeft: 20 }}
                component={RouterLink}
                to="/patientregistrationtable"
                variant="contained"
              >
                ดูข้อมูล
                            </Button>

            </Box>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}