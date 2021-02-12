//style
import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { ContentHeader ,Content, Header, Page, pageTheme,} from '@backstage/core';

//icon
import PersonIcon from '@material-ui/icons/Person';
import FaceIcon from '@material-ui/icons/Face';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
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

const Personalwelcome: FC<{}> = () => {
  const classes = useStyles();
 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="ระบบแจ้งซ่อมคอมพิวเตอร์" type="กลุ่มที่ 14">
         <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="default" href="/" startIcon={<ExitToAppIcon />}> ออกจากระบบ </Button>
     </Header>
        <Content>
          <ContentHeader title="ระบบข้อมูลบุคลากร">
            <Button 
              style={{ marginLeft: 20 }} 
              href="/Group14" 
              variant="contained"
              color="primary"
              startIcon={<CancelTwoToneIcon/>}
              > 
              ย้อนกลับ 
            </Button>
          </ContentHeader>
            <Button
              variant="outlined"
              fullWidth
              size="large"  
              href="/Personalsearch" 
              startIcon={<FaceIcon />}
              className={classes.systemstlye}> 
              ค้นหาข้อมูลบุคลากร
            </Button>
            <Button 
              variant="outlined"
              fullWidth
              size="large"
              href="/Personaltable" 
              startIcon={<PersonIcon />}
              className={classes.systemstlye}> 
              เพิ่มข้อมูลบุคลากร
            </Button>
        </Content>
   </Page>
 );
};

export default Personalwelcome;