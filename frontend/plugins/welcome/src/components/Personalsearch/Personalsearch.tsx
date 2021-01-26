//style
import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
//icon
import PersonOutlineTwoToneIcon from '@material-ui/icons/PersonOutlineTwoTone';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
//alert
import Swal from 'sweetalert2';
//entity
import { EntPersonal } from '../../api/models/EntPersonal';
//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

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
   marginLeft: '40px',
  },
  textField: {
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
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
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
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

export default function Personalsearch() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);

  const [ids, setIds] = useState(Number);

  const searchPersonals = async (id: number) => {
    const res:any = await http.getPersonal({ id: ids });
    setLoading(true);
  };

  const idshand = (event: any) => {
    setIds(event.target.value as number);
  };

  

  const personal = {
    ids : ids,
  };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

    const checkCaseSaveError = (field: string) => {
      switch(field) {
        default:
          alertMessage("error","ค้นหาข้อมูลไม่สำเร็จ");
          return;
      }
    }
  
  console.log(personal)
  function search() {
    const apiUrl = 'http://localhost:8080/api/v1/personals';
    const requestOptions = {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(personal),
    };
  
    console.log(personal); 
  
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status == true) {
          Toast.fire({
            icon: 'success',
            title: 'ค้นหาข้อมูลสำเร็จ',
  
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
      const timer = setTimeout(() => {
        setStatus(false);
      }, 1000);
    };

return (
  <Page theme={pageTheme.tool}>
    <Header
      title="ระบบข้อมูลบุคลากร" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
      <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        startIcon={<ExitToAppIcon/>}
        > 
        ออกจากระบบ 
        </Button>
    </Header>
    <Content>
      <ContentHeader title="ค้นหาข้อมูลบุคลากร">
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button  
        onClick={() => {
          search();
        }}
        variant="contained" 
        color="secondary" 
        startIcon={<SearchTwoToneIcon/>}
        > 
        ค้นหาข้อมูล 
        </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        component={RouterLink} 
        to="/Personalwelcome" 
        variant="contained"
        color="primary"
        startIcon={<CancelTwoToneIcon/>}
        > 
        ย้อนกลับ 
        </Button>
      </ContentHeader>
      <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
            
            <div className={classes.paper}><strong>ID บุคลากร</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <PersonOutlineTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="personalname"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={ids}
              onChange={idshand}
            />
          </FormControl>
        </form>
      </div>
    </Content>
   </Page>
 );
}