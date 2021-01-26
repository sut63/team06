import React, { FC , useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { Cookies } from '../../Cookie';
import { DefaultApi } from '../../api/apis'; 
import { EntDoctor } from '../../api';
import { ApiProvider } from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', 
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();
  var ck = new Cookies();
  var check : boolean
  const [path, setPath] = React.useState("");

  // list Doctor
  const [Doctor,setDoctor] = React.useState<EntDoctor[]>([])
  const listDoctor = async() => {
        const res = await api.listDoctor({})
        setDoctor(res)
  }

  // setEmail
  const [email, setEmail] = React.useState()
  const handleEmail = (event : any) => {
      setEmail(event.target.value)
  }

  // setPassword
  const [password, setPassword] = React.useState()
  const handlePassword = (event : any) => {
      setPassword(event.target.value)
  }

  // handleCookies
  function handleCookies() {
    check = ck.CheckLogin(Doctor,email,password)
    console.log("check => "+check)
    if(check === true){
      setPath("/diagnosis")
      ck.SetCookie("user_email",email,30)
      ck.SetCookie("user_id",ck.SetID(Doctor,email,password),30)
      console.log("123, "+path);
      
      //window.location.reload(false)
    }else if(check === false){
      alert(" ลงชื่อเข้าใช้ผิดพลาด !!!")
      setPath("/")
    }
  }
  // useEffect
  useEffect(() => {
      listDoctor()
  },[])

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={handleEmail}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={handlePassword}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleCookies}
            component={RouterLink}
            to={path}
          >
            Sign In
          </Button>
        </form>
      </div>
    </Container>
  );
};

export default SignIn;
