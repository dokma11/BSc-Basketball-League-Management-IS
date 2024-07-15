import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PlayerAssetCardComponent } from './player-asset-card.component';

describe('PlayerAssetCardComponent', () => {
  let component: PlayerAssetCardComponent;
  let fixture: ComponentFixture<PlayerAssetCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PlayerAssetCardComponent]
    });
    fixture = TestBed.createComponent(PlayerAssetCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
