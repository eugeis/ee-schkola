import {Component, Input} from '@angular/core';
import {PersonViewService} from '../../services/person-view.service}';

@Component({
  selector: 'app-person',
  templateUrl: './person-view.component.html',
  styleUrls: ['./person-view.component.scss'],
  providers: [PersonViewService]
})

export class PersonViewComponent {

    @Input() pageName = 'PersonComponent';
       
    constructor(public personViewService: PersonViewService) {}

}



