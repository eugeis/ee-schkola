import {Location} from '@schkola/library/LibraryApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-location',
  templateUrl: './location-basic.component.html',
  styleUrls: ['./location-basic.component.scss'],
})

export class LocationComponent {

    @Input() location: Location;
    
}




