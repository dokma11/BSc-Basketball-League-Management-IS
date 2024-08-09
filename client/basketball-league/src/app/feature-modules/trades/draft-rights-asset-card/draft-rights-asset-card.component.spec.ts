import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DraftRightsAssetCardComponent } from './draft-rights-asset-card.component';

describe('DraftRightsAssetCardComponent', () => {
  let component: DraftRightsAssetCardComponent;
  let fixture: ComponentFixture<DraftRightsAssetCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DraftRightsAssetCardComponent]
    });
    fixture = TestBed.createComponent(DraftRightsAssetCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
