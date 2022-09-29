import {Location} from '@schkola/library/LibraryApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-location',
  templateUrl: './location-basic.component.html',
  styleUrls: ['./location-basic.component.scss'],
})

export class LocationComponent implements OnInit {

    @Input() location: Location;
    
    ngOnInit() {
        if (this.location === undefined) {
            this.location = new Location();
        }
        
    }
    
}




