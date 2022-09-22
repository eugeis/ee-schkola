import {Component, Input} from '@angular/core';
import {PersonViewService} from '../../service/person-module-view.service';

@Component({
  selector: 'app-person',
  templateUrl: './person-module-view.component.html',
  styleUrls: ['./person-module-view.component.scss'],
  providers: [PersonViewService]
})

export class PersonViewComponent {

    @Input() pageName = 'PersonComponent';
       
    constructor(public personViewService: PersonViewService) {}

}



