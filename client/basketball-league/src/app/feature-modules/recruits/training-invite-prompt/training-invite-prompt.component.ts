import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, Inject, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import Map from 'ol/Map';
import View from 'ol/View';
import VectorLayer from 'ol/layer/Vector';
import Style from 'ol/style/Style';
import Icon from 'ol/style/Icon';
import OSM from 'ol/source/OSM';
import * as olProj from 'ol/proj';
import TileLayer from 'ol/layer/Tile';
import { Feature } from 'ol';
import { Point } from 'ol/geom';
import VectorSource from 'ol/source/Vector';
import { fromLonLat } from 'ol/proj';
import { InterviewInvitePromptComponent } from '../interview-invite-prompt/interview-invite-prompt.component';
import { RecruitsService } from '../recruits.service';
import { TrainingProposal, TrainingProposalStatus } from 'src/app/shared/model/trainingProposal.model';

@Component({
  selector: 'app-training-invite-prompt',
  templateUrl: './training-invite-prompt.component.html',
  styleUrls: ['./training-invite-prompt.component.css'],
  animations: [
      trigger("fadeIn", [
        transition(":enter", [
            style({ opacity: 0, transform: "translateX(-40px)" }),
            animate(
                "0.5s ease",
                style({ opacity: 1, transform: "translateX(0)" }),
            ),
        ]),
      ]),
      trigger('buttonState', [
        state('clicked', style({
          transform: 'scale(0.9)',
          opacity: 0.5
        })),
        transition('* => clicked', [
          animate('200ms')
        ]),
        transition('clicked => idle', [
          animate('200ms')
        ])
      ]),
  ],
})
export class TrainingInvitePromptComponent implements OnInit{
  buttonState: string = '';
  minDate: string;
  focused: string = '';
  public map!: Map;
  mapEnabled: boolean = false;
  newLongitude: number = 0;
  newLatitude: number = 0;
  recruitId: number = 0;
  coachId: number = 0;

  constructor(private snackBar: MatSnackBar,
              private dialog: MatDialog,
              private dialogRef: MatDialogRef<InterviewInvitePromptComponent>,
              private recruitsService: RecruitsService,
              @Inject(MAT_DIALOG_DATA) public data: any,) {
    const today = new Date();
    this.minDate = today.toISOString().split('T')[0];

    this.recruitId = data.recruitId;
    this.coachId = data.coachId;

    console.log(this.recruitId);
    console.log(this.coachId);
  }

  trainingInviteForm = new FormGroup({
    occurrenceTime: new FormControl(null, [Validators.required]),
    occurrenceDate: new FormControl(null, [Validators.required]),
    address: new FormControl('', [Validators.required]),
    longitude: new FormControl(0.0, [Validators.required]),
    latitude: new FormControl(0.0, [Validators.required]),
    duration: new FormControl('', [Validators.required]),
    selectedTrainingType: new FormControl('', [Validators.required]),
  });

  ngOnInit(): void {
    this.enableMap();    
  }

  submitInviteButtonClicked() {
    // Postavi datum i vreme
    const dateValue: Date | null = this.trainingInviteForm.value.occurrenceDate!;
    const timeValue: string | null = this.trainingInviteForm.value.occurrenceTime!;

    const [hours, minutes] = (timeValue as string).split(':');
    const dateTime = new Date(dateValue);
    dateTime.setHours(Number(hours) + 1);
    dateTime.setMinutes(Number(minutes));

    const d = new Date(dateValue);
    d.setHours(Number(hours));
    d.setMinutes(Number(minutes));

    const trainingProposal : TrainingProposal = {
      mesOdrPozTrng: this.trainingInviteForm.value.address || "",
      datVrePozTrng: dateTime || "",
      statusPozTrng: TrainingProposalStatus.PENDING,
      idTrener: this.coachId,
      idRegrut: this.recruitId,
      trajTrng: this.trainingInviteForm.value.duration?.toString() || "",
      nazTipTrng: this.trainingInviteForm.value.selectedTrainingType || ""
    }

    console.log('int lon: ' + this.trainingInviteForm.value.longitude);
    console.log('int lat: ' + this.trainingInviteForm.value.latitude);

    this.recruitsService.mapReverseSearch(this.trainingInviteForm.value.latitude!, this.trainingInviteForm.value.longitude!).subscribe(res => {
      const addressInfo = {
        number: "",
        street: "",
        city: "",
        postalCode: "",
        country: "",
      };

      let addressParts = res.display_name.split(",");

      this.setAddressInfo(addressInfo, addressParts);
        let concatenatedAddress = addressInfo.number + " " +
      addressInfo.street + " " + addressInfo.city + " " +
      addressInfo.postalCode + " " + addressInfo.country;

      trainingProposal.mesOdrPozTrng = concatenatedAddress;

      console.log(trainingProposal);

      this.recruitsService.createTrainingProposal(trainingProposal).subscribe({
        next: (result: any) => {
          this.showNotification('Training invite successfully sent!');
          this.dialogRef.close();
        }
      });
    });
  }

  enableMap(): void{
    this.mapEnabled = true;

    this.map = new Map({
        target: 'hotel_map_dialogue',
        layers: [
          new TileLayer({
            source: new OSM()
          })
        ],
        view: new View({
          center: olProj.fromLonLat([-71.06880200380476, 42.35852397178772]),
          zoom: 14
        })
    });

    this.map.on('click', (event: any) => {
        const coordinate = event.coordinate;
    
        const lonLat = olProj.toLonLat(coordinate);
    
        console.log('Longitude:', lonLat[0]);
        console.log('Latitude:', lonLat[1]);

        this.trainingInviteForm.value.longitude = lonLat[0];
        this.trainingInviteForm.value.latitude = lonLat[1];

        this.map.getLayers().forEach((layer) => {
            if (layer instanceof VectorLayer) {
              this.map.removeLayer(layer);
            }
          });

        const point = new Point(fromLonLat([lonLat[0], lonLat[1]]));

        const startMarker = new Feature(point);

        const markerStyle = new Style({
              image: new Icon({
                  anchor: [0.5, 1],
                  src: 'http://www.pngall.com/wp-content/uploads/2017/05/Map-Marker-PNG-HD-180x180.png',
                  scale: 0.4
              })
          });

        startMarker.setStyle(markerStyle);

        const vectorLayer = new VectorLayer({
              source: new VectorSource({
                  features: [startMarker]
              })
          });
        
        this.map.addLayer(vectorLayer);
    });
  }

  setAddressInfo(addressInfo: any, addressParts: any): void {
    if (addressParts.length == 10) {
        addressInfo.number = addressParts[0];
        addressInfo.street = addressParts[1];
        addressInfo.city = addressParts[4];
        addressInfo.postalCode = addressParts[8];
        addressInfo.country = addressParts[9];
    } else if (addressParts.length == 9) {
        addressInfo.number = addressParts[0];
        addressInfo.street = addressParts[1];
        addressInfo.city = addressParts[3];
        addressInfo.postalCode = addressParts[7];
        addressInfo.country = addressParts[8];
    } else if (addressParts.length == 8) {
        addressInfo.number = "";
        addressInfo.street = addressParts[1];
        addressInfo.city = addressParts[2];
        addressInfo.postalCode = addressParts[6];
        addressInfo.country = addressParts[7];
    } else if (addressParts.length == 7) {
        addressInfo.number = "";
        addressInfo.street = addressParts[0];
        addressInfo.city = addressParts[1];
        addressInfo.postalCode = addressParts[5];
        addressInfo.country = addressParts[6];
    }
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  overviewClicked(){
    this.dialogRef.close();
  }
}
