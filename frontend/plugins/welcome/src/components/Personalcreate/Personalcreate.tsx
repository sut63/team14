//style
import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
//icon
import EmailTwoToneIcon from '@material-ui/icons/EmailTwoTone';
import PersonOutlineTwoToneIcon from '@material-ui/icons/PersonOutlineTwoTone';
import VpnKeyTwoToneIcon from '@material-ui/icons/VpnKeyTwoTone';
import AddCircleTwoToneIcon from '@material-ui/icons/AddCircleTwoTone';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
//alert
import Swal from 'sweetalert2';
//entity
import { EntPersonal } from '../../api/models/EntPersonal';
import { EntGender } from '../../api/models/EntGender';
import { EntTitle } from '../../api/models/EntTitle';
import { EntDepartment } from '../../api/models/EntDepartment';

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

export default function Personalcreate() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);

  const [department, setDepartment] = useState(Number);
  const [title, setTitle] = useState(Number);
  const [gender, setGender] = useState(Number);

  const [personalname, setPersonalname] = useState(String);
  const [email, setEmail] = useState(String);
  const [password, setPassword] = useState(String);

  useEffect(() => {
    const getDepartments = async () => {
      const res = await http.listDepartment({ limit: 10, offset: 0 });
      setLoading(false);
      setDepartments(res);
      console.log(res);
    };
    getDepartments();

    const getTitles = async () => {
      const res = await http.listTitle({ limit: 10, offset: 0 });
      setLoading(false);
      setTitles(res);
      console.log(res);
    };
    getTitles();

    const getGenders = async () => {
      const res = await http.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(res);
      console.log(res);
    };
    getGenders();

  }, [loading]);

  const handlePersonalnameChange = (event: any) => {
    setPersonalname(event.target.value as string);
  };

  const handleEmailChange = (event: any) => {
    setEmail(event.target.value as string);
  };

  const handlePasswordChange = (event: any) => {
    setPassword(event.target.value as string);
  };
  
  const TitlehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTitle(event.target.value as number);
  };

  const DepartmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDepartment(event.target.value as number);
  };

  const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGender(event.target.value as number);
  };

  const personal = {
    personalname : personalname,
    email : email,
    password : password,
    title : title,
    department : department,
    gender : gender,
  };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

    const checkCaseSaveError = (field: string) => {
      switch(field) {
        case 'Personalname':
          alertMessage("error","รูปแบบชื่อไม่ถูกต้อง");
          return;
        case 'Password':
          alertMessage("error","รหัสผ่านน้อยกว่า 4 ตัว");
          return;
        case 'Email':
          alertMessage("error","รูปแบบ email ไม่ถูกต้อง");
          return;
        default:
          alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
          return;
      }
    }
  
  console.log(personal)
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/personals';
    const requestOptions = {
      method: 'POST',
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
            title: 'บันทึกข้อมูลสำเร็จ',
  
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
      <ContentHeader title="เพิ่มข้อมูลบุคลากร">
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button  
        onClick={() => {save();}}
        variant="contained" 
        color="secondary" 
        startIcon={<AddCircleTwoToneIcon/>}
        > 
        บันทึกข้อมูล 
        </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        component={RouterLink} 
        to="/Personaltable" 
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
            <div className={classes.paper}><strong>คำนำหน้าชื่อ</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="nametitle"
              value={title}
              onChange={TitlehandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกคำนำหน้าชื่อ</InputLabel>

              {titles.map((item: EntTitle) => (
                <MenuItem value={item.id}>{item.titlename}</MenuItem>
              ))}
            </Select>
            
            <div className={classes.paper}><strong>ชื่อ-นามสกุล</strong></div>
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
              value={personalname}
              onChange={handlePersonalnameChange}
            />

            <div className={classes.paper}><strong>อีเมลล์</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <EmailTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="Email"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={email}
              onChange={handleEmailChange}
            />

            <div className={classes.paper}><strong>รหัสผ่าน</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <VpnKeyTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="Password"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={password}
              onChange={handlePasswordChange}
            />
            
            <div className={classes.paper}><strong>เพศ</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="gender"
              value={gender}
              onChange={GenderhandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกเพศ</InputLabel>

              {genders.map((item: EntGender) => (
                <MenuItem value={item.id}>{item.gendername}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>แผนก</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="department"
              value={department}
              onChange={DepartmenthandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกแผนกt</InputLabel>

              {departments.map((item: EntDepartment) => (
                <MenuItem value={item.id}>{item.departmentname}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </form>
      </div>
    </Content>
   </Page>
 );
}