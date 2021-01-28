//style
import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme,} from '@backstage/core';

//icon
import PersonIcon from '@material-ui/icons/Person';
import FaceIcon from '@material-ui/icons/Face';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import ComputerTwoToneIcon from '@material-ui/icons/ComputerTwoTone';
import PhonelinkSetupTwoToneIcon from '@material-ui/icons/PhonelinkSetupTwoTone';
import NoteTwoToneIcon from '@material-ui/icons/NoteTwoTone';
import SettingsTwoToneIcon from '@material-ui/icons/SettingsTwoTone';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
    paper: {
      marginTop: theme.spacing(5),
      marginBottom: theme.spacing(5),
      marginLeft: theme.spacing(12),
      marginRight: theme.spacing(12),
    },
    systemstlye: {
      background: 'linear-gradient(45deg, #636FA4 30%, #E8CBC0 90%)',
      border: 0,
      borderRadius: 3,
      boxShadow: '0 3px 5px 2px rgba(29, 43, 100, .1)',
      color: 'white',
      height: 48,
      padding: '0 30px',
      marginRight: theme.spacing(2),
    },
  }),
);

const Group14: FC<{}> = () => {
  const classes = useStyles();
 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="ระบบแจ้งซ่อมคอมพิวเตอร์" type="กลุ่มที่ 14">
         <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="default" href="/" startIcon={<ExitToAppIcon />}> ออกจากระบบ </Button>
     </Header>

      <div className={classes.paper}>
        <Content>
            <Button 
              variant="outlined"
              fullWidth
              size="large"  
              href="/CustomerTable" 
              startIcon={<FaceIcon />}
              className={classes.systemstlye}> 
              ข้อมูลลูกค้า
            </Button>
            <Button 
              variant="outlined"
              fullWidth
              size="large"
              href="/Personalwelcome" 
              startIcon={<PersonIcon />}
              className={classes.systemstlye}> 
              ข้อมูลบุคลากร
            </Button>
            <Button 
              variant="outlined"
              fullWidth
              size="large" 
              href="Tablefix" 
              startIcon={<ComputerTwoToneIcon />}
              className={classes.systemstlye}> 
              การแจ้งซ่อม
            </Button>
            <Button 
              variant="outlined"
              fullWidth
              size="large" 
              href="/Selectadminrepair" 
              startIcon={<PhonelinkSetupTwoToneIcon />}
              className={classes.systemstlye}> 
              การซ่อม-เจ้าหน้าที่ซ่อม
            </Button>
            <Button 
              variant="outlined"
              fullWidth
              size="large" 
              href="/Tablereceipt" 
              startIcon={<NoteTwoToneIcon />}
              className={classes.systemstlye}> 
              ใบเสร็จ
            </Button>
            <Button 
              variant="outlined"
              fullWidth
              size="large" 
              href="/Producttables" 
              startIcon={<SettingsTwoToneIcon />}
              className={classes.systemstlye}> 
              ข้อมูลอะไหล่
            </Button>
        </Content>
      </div>
   </Page>
 );
};

export default Group14;