import { Component, OnInit, Input } from '@angular/core';
import { Sock } from 'src/app/dataModel/sock.model';
import { SocksManagerService } from '../services/socksManager/socks-manager.service';
import { Observable } from 'rxjs';
import {
  MatchService,
  OtherInfo,
} from '../services/match/match-service.service';
import { ActivatedRoute, Params } from '@angular/router';

@Component({
  selector: 'app-match-result',
  templateUrl: './match-result.component.html',
  styleUrls: ['./match-result.component.scss'],
})
export class MatchResultComponent implements OnInit {
  mySock!: Observable<Sock>;
  otherInfo!: Observable<OtherInfo>;

  @Input()
  result = '';

  other = {
    firstname: 'Eliott',
    surname: 'Chytil',
    address: "Grand'Rue 3",
    city: 'Begnins',
    npa: '1268',
    country: 'Switzerland',
  };

  sentence!: string;
  heartIcon!: string;

  constructor(
    private route: ActivatedRoute,
    private sockManager: SocksManagerService,
    private matchService: MatchService
  ) {}

  ngOnInit(): void {

    this.route.params.subscribe({
      next: (params: any) => {
        this.mySock = this.matchService.getSelfSock_id(params.id);
        this.otherInfo = this.matchService.getOtherInfos_id(params.id);
      },
    });

    if (this.result === 'win') {
      this.sentence = 'You will receive the sock in a few days';
      this.heartIcon = 'favorite';
    } else {
      this.sentence = 'Please send your sock at the following address:';
      this.heartIcon = 'heart_broken';

      this.other.firstname = 'Eliott';
      this.other.surname = 'Chytil';
      this.other.address = "Grand'Rue 3";
      this.other.city = 'Begnins';
      this.other.npa = '1268';
      this.other.country = 'Switzerland';
    }
  }
}
