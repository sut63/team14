import React, { FC, useEffect,useState } from 'react';
import {Link as RouterLink} from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import { ContentHeader,Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
}from '@material-ui/core';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import SaveIcon from '@material-ui/icons/Save';
import { DefaultApi } from '../../api/apis';
import { EntPersonal } from '../../api/models/EntPersonal';
import { EntFix } from '../../api/models/EntFix';
import { EntProduct } from '../../api/models/EntProduct';
import { EntAdminrepair } from '../../api/models/EntAdminrepair';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import { Alert } from '@material-ui/lab';
import Button from '@material-ui/core/Button';
import AdminrepairTable from '../AdminrepairTable';
import AdminrepairFixTable from '../AdminrepairFixTable';
import AdminrepairProductTable from '../AdminrepairProductTable';
import { Cookies } from '../WelcomePage/Cookie'

const HeaderCustom = {
  minHeight: '50px',
};

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 20,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
    marginRight: theme.spacing(-1),
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
}));

  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

const Adminrepair: FC<{}> = () =>{
  //const UI = { giveName : 'Confirmation'}
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);
  const [adtable, setAdtable] = useState(false);
  const [fixtable, setFixtable] = useState(false);
  const [producttable, setProducttable] = useState(false);

  const [personals, setPersonals] = useState<EntPersonal[]>([]);
  const [fixs, setFixs] = useState<EntFix[]>([]);
  const [products, setProducts] = useState<EntProduct[]>([]);

  const [fixid, setFixID] = useState(Number);
  const [equipmentdamate, setEquipmentdamate] = useState(String);
  const [productid, setProductID] = useState(Number);

  useEffect(() =>{
    const getFixs = async () =>{
        const res = await api.listFix({limit:10,offset:0});
        setLoading(false);
        setFixs(res);
      }
      getFixs();

    const getProducts = async () =>{
        const res = await api.listProduct({limit:10,offset:0});
        setLoading(false);
        setProducts(res);
    }
    getProducts();
    
  },[loading]);

  
  const Equipmentdamatehandlechange = (event: any) => {
    setEquipmentdamate(event.target.value as string);
  }

  const Fixhandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFixID(event.target.value as number);
  }

  const Producthandlechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProductID(event.target.value as number);
  }

  const CheckTableAdminrepair = (event: any) => {
    setAdtable(true)
    setFixtable(false)
    setProducttable(false)
  }
  const CheckTableFix = (event: any) => {
    setAdtable(false)
    setFixtable(true)
    setProducttable(false)
  }
  const CheckTableProduct = (event: any) => {
    setAdtable(false)
    setFixtable(false)
    setProducttable(true)
  }
  
  const CreateAdminrepair = async ()=>{
    if ((fixid != null)  && (equipmentdamate != null) && (equipmentdamate != "") && (productid != null) ){
    const Adminrepair = {
      personal :  Number(cookieID),
      fix : fixid,
      equipmentdamate : equipmentdamate,
      product : productid,
    };
    console.log(Adminrepair);
    setStatus(true);
    setAlert(false);
    const res: any = await api.createAdminrepair({ adminrepair : Adminrepair});
    setAlert(true);
      if (res.id != '') {
        const timer = setTimeout(() => {
          window.location.reload(false);
       }, 3000);
      }
    }else {
      setStatus(true);
      setAlert(false);
    }
    
  }

  return (
    <Page theme={pageTheme.home}>
    <Header title="ระบบบันทึกซ่อมแซมคอมพิวเตอร์ของพนักงาน" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
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
    {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกข้อมูลเสร็จสิ้น!!!
               </Alert>
             ) : (
               <Alert severity="error" style={{ marginTop: 20,marginBottom: 30 }}>
                 มีข้อมูลบางอย่างไม่ถูกต้อง กรุณากรอกข้อมูลใหม่ อีกครั้ง!!!
               </Alert>
             )}
           </div>
         ) : null}
         <ContentHeader title="เพิ่มข้อมูลบันทึกซ่อมแซมคอมพิวเตอร์ของ พนักงาน">
    <Button
      style={{ marginLeft: 10}}
      variant="outlined"
      color="primary"
      onClick={CheckTableFix}
      >
        ดู บันทึกแจ้งซ่อมแซมคอมพิวเตอร์
    </Button>
    <Button
      style={{ marginLeft: 10 }}
      variant="outlined"
      onClick={CheckTableProduct}
      >
        ดู รายละเอียดอะไหล่คอมพิวเตอร์
    </Button>
    <Button
      style={{ marginLeft: 10 }}
      variant="outlined"
      color="secondary"
      onClick={CheckTableAdminrepair}
      >
        ดู บันทึกการซ่อมแซมคอมพิวเตอร์ของพนักงาน
    </Button>
    </ContentHeader>
    {adtable ?(<AdminrepairTable></AdminrepairTable>) : null}
    {producttable ?(<AdminrepairProductTable></AdminrepairProductTable>) : null}
    {fixtable ?(<AdminrepairFixTable></AdminrepairFixTable>) : null}
       
      <Container maxWidth="sm">
        <Grid container spacing={3}>
          <Grid item xs={12}></Grid>
          <Grid item xs={3}>
            <div className={classes.paper}>หมายเลข บันทึกแจ้งซ่อมแซมคอมพิวเตอร์</div>
          </Grid>
          <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>เลือกบันทึกแจ้งซ่อมแซมคอมพิวเตอร์</InputLabel>
              <Select
                name="fix"
                value={ fixid }
                onChange={Fixhandlechange}
              >
                {fixs.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.id}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <div className={classes.paper}>ความเสียหาย</div>
          </Grid>
          <Grid item xs={9}>
            <TextField 
            style={{ width: 300}}
            id="equipmentdamate" 
            label="กรุณากรอกข้อมูลความเสียหาย" 
            variant="standard"
            type="string"
            size="medium"
            value={equipmentdamate}
            onChange = {Equipmentdamatehandlechange}/>
          </Grid>

          <Grid item xs={3}>
            <div className={classes.paper}>อะไหล่คอมพิวเตอร์</div>
          </Grid>
          <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>เลือกอะไหล่คอมพิวเตอร์</InputLabel>
              <Select
                name="product"
                value={ productid }
                onChange={Producthandlechange}
              >
                {products.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.productname}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={3}></Grid>
          <Grid item xs={12}>
            <Button
              style={{ marginRight: 30}}
              variant="contained"
              color="primary"
              size="large"
              onClick={CreateAdminrepair}
              startIcon={<SaveIcon/>}
            >
              บันทึกข้อมูลการซ่อมแซม
            </Button>
            <Button style={{ marginRight: 30}} 
            variant="contained" color="secondary" href="/Group14" startIcon={<HomeRoundedIcon/>}> กลับหน้าหลัก </Button>
          </Grid>
        </Grid>
      </Container>
    </Content>
  </Page>
  );
};
export default Adminrepair;