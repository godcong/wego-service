import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';
import {ActivatedRoute, ParamMap} from '@angular/router';
import {DataService} from '../data.service';
import {WechatService} from '../wechat.service';
import {Location} from '@angular/common';

@Component({
  selector: 'app-qrcode',
  templateUrl: './qrcode.component.html',
  styleUrls: ['./qrcode.component.scss']
})
export class QrcodeComponent implements OnInit {
  public size: SizeService;
  public qrcode: string;
  private router: ActivatedRoute;
  private id: string;
  private user: string;
  private data: DataService;
  private wechat: WechatService;
  private url: string;
  private userInfo: any;

  constructor(size: SizeService, router: ActivatedRoute, data: DataService, wechat: WechatService, loc: Location) {
    console.log(loc.path(true));
    this.size = size;
    this.router = router;
    this.data = data;
    this.wechat = wechat;
    this.qrcode = 'http://localhost:8080';
    this.url = loc.path(true);
  }

  ngOnInit() {
    this.router.queryParamMap.subscribe((params: ParamMap) => {
      console.log(params);
      this.user = params.get('user');
    });
    this.router.paramMap.subscribe((params: ParamMap) => {
      console.log(params);
      this.id = params.get('id');
    });

    this.data.getMyInfo().subscribe((params: any) => {
      console.log(params);
      this.userInfo = params;
    }, (error: any) => {
      console.log(error);
      alert(error.error.message);
    });

    this.data.getSpreadShareInfo(this.id, this.user, this.url).subscribe((params: any) => {
      console.log(params);
      this.qrcode = params.url;
      this.wechat.init(params.config);
      this.wechat.shareAppMessage(this.qrcode, this.userInfo.headimgurl);
    }, (error: any) => {
      console.log(error);
      alert(error.error.message);
    });
  }

  onSuccess() {
    console.log(this.qrcode);
  }
}
