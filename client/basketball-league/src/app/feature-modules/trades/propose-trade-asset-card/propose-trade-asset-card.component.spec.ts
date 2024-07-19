import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProposeTradeAssetCardComponent } from './propose-trade-asset-card.component';

describe('ProposeTradeAssetCardComponent', () => {
  let component: ProposeTradeAssetCardComponent;
  let fixture: ComponentFixture<ProposeTradeAssetCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ProposeTradeAssetCardComponent]
    });
    fixture = TestBed.createComponent(ProposeTradeAssetCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
