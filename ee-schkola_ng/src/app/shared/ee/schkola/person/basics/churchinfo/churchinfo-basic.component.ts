import {ChurchInfo} from '@schkola/person/PersonApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-churchinfo',
  templateUrl: './churchinfo-basic.component.html',
  styleUrls: ['./churchinfo-basic.component.scss'],
})

export class ChurchInfoComponent {

    @Input() churchinfo: ChurchInfo;
    
}




