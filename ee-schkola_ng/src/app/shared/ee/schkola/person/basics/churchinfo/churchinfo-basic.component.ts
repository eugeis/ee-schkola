import {ChurchInfo} from '@schkola/person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-churchinfo',
  templateUrl: './churchinfo-basic.component.html',
  styleUrls: ['./churchinfo-basic.component.scss'],
})

export class ChurchInfoComponent implements OnInit {

    @Input() churchinfo: ChurchInfo;
    
    ngOnInit() {
        if (this.churchinfo === undefined) {
            this.churchinfo = new ChurchInfo();
        }
        
    }
    
}




